package main

import (
    "reflect"
    "testing"
)

func TestWordCount(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected map[string]int
    }{
        {
            name:  "basic word counting",
            input: "Hello, hello! How are you?",
            expected: map[string]int{
                "hello": 2,
                "how":   1,
                "are":   1,
                "you":   1,
            },
        },
        {
            name:  "empty string",
            input: "",
            expected: map[string]int{},
        },
        {
            name:  "string with only punctuation",
            input: "!!!,,,???",
            expected: map[string]int{},
        },
        {
            name:  "mixed case words",
            input: "Hello HELLO hello",
            expected: map[string]int{
                "hello": 3,
            },
        },
        {
            name:  "words with punctuation",
            input: "Hello! How are you?",
            expected: map[string]int{
                "hello": 1,
                "how":   1,
                "are":   1,
                "you":   1,
            },
        },
        {
            name:  "multiple spaces between words",
            input: "Hello   World",
            expected: map[string]int{
                "hello": 1,
                "world": 1,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := WordCount(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("WordCount(%q) = %v, want %v", tt.input, result, tt.expected)
            }
        })
    }
} 