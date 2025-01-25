package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ConnectToRedis establishes a connection to a Redis server located at the specified host.
// It returns a pointer to the Redis client if the connection is successful, or an error if the connection fails.
func ConnectToRedis(host string) (*redis.Client, error) {
	// Create a new Redis client with the provided host address and default configuration.
	rdb := redis.NewClient(&redis.Options{
		Addr:     host, // Address of the Redis server.
		Password: "",   // No password is set for the connection.
		DB:       0,    // Use the default Redis database (DB 0).
	})

	// Ping the Redis server to verify the connection.
	c := rdb.Ping(context.TODO())
	if c.Err() != nil {
		// If there's an error during the ping, return the error.
		return nil, c.Err()
	}

	// If the connection is successful, return the Redis client.
	return rdb, nil
}
