package assert

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}, msg ...string) {
    if !reflect.DeepEqual(expected, actual) {
        msglToWPrint := fmt.Sprint(msg)
        if len(msg) == 0 {
            msglToWPrint = fmt.Sprintf("expected %v (type %T), but got %v (type %T)", expected, expected, actual, actual)
        } 
        t.Fatal(msglToWPrint)
    }
  
}

func Contains(t *testing.T, s, expectedSubstr string, msg ...string) {
    if !strings.Contains(s, expectedSubstr) {
        msglToWPrint := fmt.Sprint(msg)
        if len(msg) == 0 {
            msglToWPrint = fmt.Sprintf("expected %q to contain %q", s, expectedSubstr)
        }
        t.Fatal(msglToWPrint)
    } 
}

func NotContains(t *testing.T, s, unexpectedSubstr string, msg ...string) {
    if strings.Contains(s, unexpectedSubstr) {
        msglToWPrint := fmt.Sprint(msg)
        if len(msg) == 0 {
            msglToWPrint = fmt.Sprintf("expected %q to not contain %q", s, unexpectedSubstr)
        }
        t.Fatal(msglToWPrint)
    } 
}

func IsError(t *testing.T, err error, msg ...string) {
    if err == nil {
        msglToWPrint := fmt.Sprint(msg)
        if len(msg) == 0 {
            msglToWPrint = "expected an error, but got nil"
        }
        t.Fatal(msglToWPrint)
    }
}

func NoError(t *testing.T, err error, msg ...string) {
    if err != nil {
        msglToWPrint := fmt.Sprint(msg)
        if len(msg) == 0 {
            msglToWPrint = fmt.Sprintf("expected no error, but got: %v", err)
        }
        t.Fatal(msglToWPrint)
    }
}

