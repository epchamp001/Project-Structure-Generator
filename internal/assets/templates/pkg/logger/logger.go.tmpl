// Package logger provides a flexible, functional-options-based wrapper around Uber's zap logger.
//
// Recommended order of options when calling NewLogger:
// 1. Profile options (sets base config):
//    - WithMode("dev" / "production")
// 2. Core settings (override profile defaults):
//    - WithLevel(zapcore.DebugLevel / InfoLevel / WarnLevel / ErrorLevel / DPanicLevel / PanicLevel / FatalLevel)
//    - WithEncoding("console" / "json")
//    - WithSampling(&zap.SamplingConfig{Initial: int, Thereafter: int}) or WithSampling(nil)
// 3. Output and fields:
//    - WithOutputPaths("stdout", "/path/to/file.log")
//    - WithErrorOutputPaths("stderr", "/path/to/error.log")
//    - WithInitialFields(map[string]interface{}{"service": "my-service"})
// 4. Fine-tuning encoder and metadata:
//    - WithDisableCaller(true / false)
//    - WithDisableStacktrace(true / false)
//    - WithEncoderConfig(func(ec *zapcore.EncoderConfig) { ... })
//    - WithEncoderConfig examples:
//        ec.TimeKey = "timestamp"
//        ec.EncodeTime = zapcore.ISO8601TimeEncoder
//        ec.EncodeLevel = zapcore.CapitalColorLevelEncoder  // for colored levels
//
// Example usage:
// log, err := logger.NewLogger(
//     // 1) Profile
//     logger.WithMode("dev"),
//     // 2) Core settings
//     logger.WithLevel(zapcore.DebugLevel),
//     logger.WithEncoding("console"),
//     logger.WithSampling(&zap.SamplingConfig{Initial:50, Thereafter:100}),
//     // 3) Output and fields
//     logger.WithOutputPaths("stdout", "/var/log/app.log"),
//     logger.WithErrorOutputPaths("stderr"),
//     logger.WithInitialFields(map[string]interface{}{"service":"auth"}),
//     // 4) Fine-tuning
//     logger.WithDisableCaller(false),
//     logger.WithDisableStacktrace(false),
//     logger.WithEncoderConfig(func(ec *zapcore.EncoderConfig) {
//         ec.TimeKey = "ts"
//         ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
//     }),
// )
// if err != nil {
//     panic(err)
// }
// defer log.Sync()

package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

type ZapLogger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

type Option func(cfg *zap.Config)

// WithMode sets the logging mode on the zap.Config.
func WithMode(mode string) Option {
	return func(cfg *zap.Config) {
		switch mode {
		case "dev", "development":
			*cfg = zap.NewDevelopmentConfig()
			cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		case "prod", "production":
			*cfg = zap.NewProductionConfig()
			cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

		default:
			fmt.Fprintf(os.Stderr,
				"[logger] warning: unknown mode %q, defaulting to production\n",
				mode,
			)
		}

		cfg.EncoderConfig.TimeKey = "timestamp"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}
}

// WithLevel sets the minimum logging level
func WithLevel(l zapcore.Level) Option {
	return func(cfg *zap.Config) {
		cfg.Level = zap.NewAtomicLevelAt(l)
	}
}

// WithDisableCaller enables or disables caller annotation
func WithDisableCaller(disable bool) Option {
	return func(cfg *zap.Config) {
		cfg.DisableCaller = disable
		if disable {
			cfg.EncoderConfig.CallerKey = ""
		} else if cfg.EncoderConfig.CallerKey == "" {
			cfg.EncoderConfig.CallerKey = "caller"
		}
	}
}

// WithDisableStacktrace enables or disables automatic stacktrace collection
func WithDisableStacktrace(disable bool) Option {
	return func(cfg *zap.Config) {
		cfg.DisableStacktrace = disable
	}
}

// WithSampling sets the sampling policy
func WithSampling(s *zap.SamplingConfig) Option {
	return func(cfg *zap.Config) {
		cfg.Sampling = s
	}
}

// WithEncoding selects the encoder: "console" or "json"
func WithEncoding(enc string) Option {
	return func(cfg *zap.Config) {
		cfg.Encoding = enc
	}
}

// WithEncoderConfig provides full access to zapcore.EncoderConfig fields
func WithEncoderConfig(f func(ec *zapcore.EncoderConfig)) Option {
	return func(cfg *zap.Config) {
		f(&cfg.EncoderConfig)
	}
}

// WithOutputPaths specifies where to write logs (stdout, files, etc.)
func WithOutputPaths(paths ...string) Option {
	return func(cfg *zap.Config) {
		cfg.OutputPaths = paths
	}
}

// WithErrorOutputPaths specifies where to write internal logger errors
func WithErrorOutputPaths(paths ...string) Option {
	return func(cfg *zap.Config) {
		cfg.ErrorOutputPaths = paths
	}
}

// WithInitialFields adds fields to every log entry
func WithInitialFields(fields map[string]interface{}) Option {
	return func(cfg *zap.Config) {
		cfg.InitialFields = fields
	}
}

// NewLogger assembles the zap.Config and constructs the Logger
func NewLogger(opts ...Option) (Logger, error) {
	cfg := zap.NewProductionConfig()

	for _, opt := range opts {
		opt(&cfg)
	}

	z, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger: z,
		sugar:  z.Sugar(),
	}, nil
}

func (l *ZapLogger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *ZapLogger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *ZapLogger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *ZapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

func (l *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

func (l *ZapLogger) Sync() {
	_ = l.logger.Sync()
}
