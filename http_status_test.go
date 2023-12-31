package main

import (
    "testing"
)

func TestHTTPStatusString(t *testing.T) {
    tests := []struct {
        status HTTPStatus
        want   string
    }{
        {StatusOK, "OK"},
        {StatusUnauthorized, "Unauthorized"},
        {StatusPaymentRequired, "Payment Required"},
        {StatusForbidden, "Forbidden"},
        {999, "HTTPStatus(999)"}, // 未定義のステータスコードのテスト
    }

    for _, tt := range tests {
        got := tt.status.String()
        if got != tt.want {
            t.Errorf("HTTPStatus.String() = %v, want %v", got, tt.want)
        }
    }
}
