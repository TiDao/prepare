#ifndef PAILLIER_PAILLIER_EXTENSION_H
#define PAILLIER_PAILLIER_EXTENSION_H

#include "./paillier.h"

void
paillier_key_gen( int modulusbits,
                  paillier_pubkey_t** pubkey,
                  paillier_prvkey_t** prvke,
                  paillier_get_rand_t get_rand );

char*
paillier_encrypt( paillier_pubkey_t *pubkey,
                  char* pt,
                  paillier_get_rand_t get_rand);

char*
paillier_decrypt( paillier_pubkey_t* pubkey,
                  paillier_prvkey_t* prvkey,
                  char* ct_ptr );

char*
paillier_add_cipher( paillier_pubkey_t *pubkey,
                     char* ct0_ptr,
                     char* ct1_ptr );

char*
paillier_add_plain( paillier_pubkey_t *pubkey,
                    char* ct_ptr,
                    char* pt );

char*
paillier_sub_cipher( paillier_pubkey_t *pubkey,
                     char* ct0_ptr,
                     char* ct1_ptr );

char*
paillier_sub_plain( paillier_pubkey_t *pubkey,
                    char* ct,
                    char* pt );

char*
paillier_num_mul( paillier_pubkey_t *pubkey,
                  char* ct_ptr,
                  char* pt );

void
paillier_reverse(mpz_t res, mpz_t a, mpz_t n);

void
adjust_encrypt_plaintext_domain(paillier_pubkey_t* pub,
                                paillier_plaintext_t* pt);

void
adjust_decrypt_plaintext_domain(paillier_pubkey_t* pub,
                                paillier_plaintext_t* pt);

#endif //PAILLIER_PAILLIER_EXTENSION_H
