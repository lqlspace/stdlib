##
[链接](https://www.tutorialspoint.com/cryptography/message_authentication.htm)


## Stream cipher和Block cipher

## 如果第三方偷换 公钥给sender，怎么办？选择PKI

## PKCS#5和PKCS7
pkcs是""Public Key Cryptography Standards"的简称，pkcs#5是
pkcs#7的一种特殊情况，即blockSize = 8byte，而pkcs7则可以是小于255bytes的任意自己；
AES的block size是16bytes，所以AES采用的一定是pkcs7

## AES CBC生成的密文和明文长度一样？
一般stream cipher算法每bit或byte逐个进行加密，密文和明文长度保持一致；block cipher
算法以blocksize为单位进行加密，加密后的长度比明文多padding个长度；

## 当明文长度是blockSize的整数倍时，是否仍需要padding?
一般都加上，即添加一个blockSize大小的padding，值为blockSize;
所以padding的取值范围一个是(1,blockSize]

## 有关GCM的问题回答链接
[AES-GCM](https://crypto.stackexchange.com/questions/17999/aes256-gcm-can-someone-explain-how-to-use-it-securely-ruby?answertab=votes#tab-top)
