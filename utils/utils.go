package utils

import (
    "crypto/rand"
    "math/big"
    "strings"
)

func GenerateRandomUUID() (string, error) {
    const length = 16
    const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

    var sb strings.Builder
    sb.WriteString("uuid-")

    for i := 0; i < length; i++ {
        idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
        if err != nil {
            return "", err
        }
        sb.WriteByte(chars[idx.Int64()])
    }

    return sb.String(), nil
}

