package runner

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want *Runner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunner_Add(t *testing.T) {
	type fields struct {
		interrupt chan os.Signal
		complete  chan error
		timeout   <-chan time.Time
		tasks     []func(int)
	}
	type args struct {
		tasks []func(int)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{
				interrupt: tt.fields.interrupt,
				complete:  tt.fields.complete,
				timeout:   tt.fields.timeout,
				tasks:     tt.fields.tasks,
			}
			r.Add(tt.args.tasks...)
		})
	}
}

func TestRunner_Start(t *testing.T) {
	type fields struct {
		interrupt chan os.Signal
		complete  chan error
		timeout   <-chan time.Time
		tasks     []func(int)
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{
				interrupt: tt.fields.interrupt,
				complete:  tt.fields.complete,
				timeout:   tt.fields.timeout,
				tasks:     tt.fields.tasks,
			}
			if err := r.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Runner.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunner_run(t *testing.T) {
	type fields struct {
		interrupt chan os.Signal
		complete  chan error
		timeout   <-chan time.Time
		tasks     []func(int)
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{
				interrupt: tt.fields.interrupt,
				complete:  tt.fields.complete,
				timeout:   tt.fields.timeout,
				tasks:     tt.fields.tasks,
			}
			if err := r.run(); (err != nil) != tt.wantErr {
				t.Errorf("Runner.run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunner_gotInterrupt(t *testing.T) {
	type fields struct {
		interrupt chan os.Signal
		complete  chan error
		timeout   <-chan time.Time
		tasks     []func(int)
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{
				interrupt: tt.fields.interrupt,
				complete:  tt.fields.complete,
				timeout:   tt.fields.timeout,
				tasks:     tt.fields.tasks,
			}
			if got := r.gotInterrupt(); got != tt.want {
				t.Errorf("Runner.gotInterrupt() = %v, want %v", got, tt.want)
			}
		})
	}
}
