package minio

import (
	"flag"
	"reflect"
	"testing"

	"github.com/minio/cli"
)

func TestMinioContextEnv(t *testing.T) {
	t.Setenv("ENDPOINT", "s3.amazonaws.com")
	t.Setenv("ACCESS_KEY", "test-access")
	t.Setenv("SECRET_KEY", "test-secret")
	t.Setenv("TRACE", "1")
	t.Setenv("SECURE", "1")

	set := flag.NewFlagSet("test", flag.ContinueOnError)
	cliCtx := cli.NewContext(nil, set, nil)

	ctx := newMinioContext(cliCtx)
	expected := minioContext{
		Endpoint:  "s3.amazonaws.com",
		AccessKey: "test-access",
		SecretKey: "test-secret",
		Trace:     true,
		Secure:    true,
	}

	if !reflect.DeepEqual(ctx, expected) {
		t.Errorf("got %+v, expected %+v", ctx, expected)
	}
}
