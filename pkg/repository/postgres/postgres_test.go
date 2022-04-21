//go:build integration
// +build integration

package postgres_test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"log"
	"net"
	"net/url"
	"path/filepath"
	"runtime"
	"time"
)

func startPG() (*sql.DB, func() error) {
	pgURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("myuser", "mypass"),
		Path:   "mydatabase",
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	pw, _ := pgURL.User.Password()
	env := []string{
		"POSTGRES_USER=" + pgURL.User.Username(),
		"POSTGRES_PASSWORD=" + pw,
		"POSTGRES_DB=" + pgURL.Path,
	}
	abs, err := filepath.Abs("../../../db")
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{Repository: "postgres", Cmd: []string{"postgres", "-c", "shared_preload_libraries=pg_stat_statements"}, Tag: "14-alpine", Env: env}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.Mounts = []docker.HostMount{
			{
				Target: "/docker-entrypoint-initdb.d",
				Source: abs,
				Type:   "bind",
			},
		}
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start postgres container: %v", err)
	}

	pgURL.Host = resource.Container.NetworkSettings.IPAddress

	// Docker layer network is different on Mac
	if runtime.GOOS == "darwin" {
		pgURL.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}
	var db *sql.DB
	pool.MaxWait = 10 * time.Second
	err = pool.Retry(func() (err error) {
		db, err = sql.Open("postgres", pgURL.String())
		if err != nil {
			return err
		}
		return db.Ping()
	})
	if err != nil {
		log.Fatal(err)
	}
	return db, func() error {
		return pool.Purge(resource)
	}
}
