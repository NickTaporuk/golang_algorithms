package queue

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Queue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type fields struct {
		start  *node
		end    *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Queue{
				start:  tt.fields.start,
				end:    tt.fields.end,
				length: tt.fields.length,
			}
			if got := this.Dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type fields struct {
		start  *node
		end    *node
		length int
	}
	type args struct {
		value interface{}
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
			this := &Queue{
				start:  tt.fields.start,
				end:    tt.fields.end,
				length: tt.fields.length,
			}
			this.Enqueue(tt.args.value)
		})
	}
}

func TestQueue_Len(t *testing.T) {
	type fields struct {
		start  *node
		end    *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Queue{
				start:  tt.fields.start,
				end:    tt.fields.end,
				length: tt.fields.length,
			}
			if got := this.Len(); got != tt.want {
				t.Errorf("Queue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type fields struct {
		start  *node
		end    *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Queue{
				start:  tt.fields.start,
				end:    tt.fields.end,
				length: tt.fields.length,
			}
			if got := this.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
