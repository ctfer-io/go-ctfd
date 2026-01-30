package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ctfer-io/go-ctfd/api"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

const (
	url = "http://localhost:8000"
)

func main() {
	ctx := context.Background()
	trans := otelhttp.NewTransport(http.DefaultTransport)

	// Setup OpenTelemetry tracer and exporter
	tp, err := initTracer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	tr := otel.Tracer("example/opentelemetry")

	// Then create an overall span for the whole code that follows
	ctx, span := tr.Start(ctx, "example-code")
	defer span.End()

	// Connect to CTFd
	fmt.Println("[+] Getting initial nonce and session values")
	nonce, session, err := func(ctx context.Context) (string, string, error) {
		ctx, span := tr.Start(ctx, "get-nonce-and-session")
		defer span.End()

		return api.GetNonceAndSession(url, api.WithContext(ctx), api.WithTransport(trans))
	}(ctx)
	if err != nil {
		log.Fatalf("Getting nonce and session: %s", err)
	}
	cli := api.NewClient(url, nonce, session, "")

	// Then setup CTFd
	fmt.Println("[+] Setting up CTFd")
	err = func(ctx context.Context) error {
		ctx, span := tr.Start(ctx, "setup")
		defer span.End()

		return cli.Setup(&api.SetupParams{
			CTFName:                "24h IUT",
			CTFDescription:         "24h IUT annual Cybersecurity CTF.",
			UserMode:               "users",
			Name:                   "PandatiX",
			Email:                  "lucastesson@protonmail.com",
			Password:               "password",
			ChallengeVisibility:    "public",
			AccountVisibility:      "public",
			ScoreVisibility:        "public",
			RegistrationVisibility: "public",
			VerifyEmails:           false,
			TeamSize:               nil,
			CTFLogo:                nil,
			CTFBanner:              nil,
			CTFSmallIcon:           nil,
			CTFTheme:               "core",
			ThemeColor:             "",
			Start:                  "",
			End:                    "",
		}, api.WithContext(ctx), api.WithTransport(trans))
	}(ctx)
	if err != nil {
		log.Fatalf("Setting up CTFd: %s", err)
	}

	// Create an API key
	fmt.Println("[+] Creating API Token")
	token, err := func(ctx context.Context) (*api.Token, error) {
		ctx, span := tr.Start(ctx, "create-api-token")
		defer span.End()

		return cli.PostTokens(&api.PostTokensParams{
			Expiration:  "2222-02-02",
			Description: "Example API token.",
		}, api.WithContext(ctx), api.WithTransport(trans))
	}(ctx)
	if err != nil {
		log.Fatalf("Creating API token: %s", err)
	}
	cli.SetAPIKey(*token.Value)

	// And finally create a challenge
	fmt.Println("[+] Creating challenge")
	err = func(ctx context.Context) error {
		ctx, span := tr.Start(ctx, "create-challenge")
		defer span.End()

		_, err := cli.PostChallenges(&api.PostChallengesParams{
			Name:           "Break The License 1/2",
			Category:       "crypto",
			Description:    "...",
			Attribution:    ptr("pandatix"),
			Function:       ptr("logarithmic"),
			ConnectionInfo: ptr("ssh -l user@crypto1.ctfer.io"),
			MaxAttempts:    ptr(3),
			Initial:        ptr(500),
			Decay:          ptr(17),
			Minimum:        ptr(50),
			State:          "visible",
			Type:           "dynamic",
		}, api.WithContext(ctx), api.WithTransport(trans))
		return err
	}(ctx)
	if err != nil {
		log.Fatalf("Creating challenge: %s", err)
	}

	fmt.Printf("Waiting for few seconds to export spans ...\n\n")
	time.Sleep(10 * time.Second)
}

func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("example"),
		),
	)
	if err != nil {
		return nil, err
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(r),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, err
}

func ptr[T any](t T) *T {
	return &t
}
