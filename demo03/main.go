package main

import (
	"fmt"
	"math/big"
)

type RSAKey struct {
	N *big.Int
	E *big.Int
	D *big.Int
}

func modMultiply(a, b, m *big.Int) *big.Int {
	result := new(big.Int).SetUint64(0)
	one := new(big.Int).SetUint64(1)

	for b.Cmp(one) > 0 {
		if new(big.Int).And(b, one).Cmp(one) == 0 {
			result.Add(result, a)
			result.Mod(result, m)
		}
		a.Lsh(a, 1)
		a.Mod(a, m)
		b.Rsh(b, 1)
	}

	return result
}

func modExp(base, exp, mod *big.Int) *big.Int {
	result := new(big.Int).SetUint64(1)
	one := new(big.Int).SetUint64(1)

	for exp.Cmp(one) > 0 {
		if new(big.Int).And(exp, one).Cmp(one) == 0 {
			result.Mul(result, base)
			result.Mod(result, mod)
		}
		base.Mul(base, base)
		base.Mod(base, mod)
		exp.Rsh(exp, 1)
	}

	return result
}

func encrypt(plaintext *big.Int, publicKey RSAKey) *big.Int {
	return modExp(plaintext, publicKey.E, publicKey.N)
}

func decrypt(ciphertext *big.Int, privateKey RSAKey) *big.Int {
	return modExp(ciphertext, privateKey.D, privateKey.N)
}

func main() {
	var publicKey RSAKey
	var privateKey RSAKey

	// 输入公钥和私钥参数
	fmt.Println("请输入公钥 (N E):")
	publicKey.N = new(big.Int)
	publicKey.E = new(big.Int)
	fmt.Scan(publicKey.N)
	fmt.Scan(publicKey.E)

	fmt.Println("请输入私钥 (N D):")
	privateKey.N = new(big.Int)
	privateKey.D = new(big.Int)
	fmt.Scan(privateKey.N)
	fmt.Scan(privateKey.D)

	// 输入要加密的明文
	fmt.Print("请输入要加密的明文: ")
	plaintext := new(big.Int)
	plaintextStr := ""
	fmt.Scan(&plaintextStr, plaintext)
	plaintext.SetString(plaintextStr, 10)

	// 加密明文
	ciphertext := encrypt(plaintext, publicKey)
	fmt.Println("加密后的密文:", ciphertext)

	// 解密密文
	decrypted := decrypt(ciphertext, privateKey)
	fmt.Println("解密后的明文:", decrypted)
}
