#!/usr/bin/env bash

SIMD_BIN=${SIMD_BIN:=$(which simd 2>/dev/null)}

ALICE_MNEMONIC="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art"
BOB_MNEMONIC="claim infant gather cereal sentence general cheese float hero dwarf miracle oven tide virus question choice say relax similar rice surround deal smooth rival"
CHARLIE_MNEMONIC="letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless"
DAVID_MNEMONIC="zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote"
UNKNOWN_MNOMONIC="purpose clutch ill track skate syrup cost among piano elegant close chaos come quit orchard acquire plunge hockey swift tongue salt supreme sting night"

if [ -z "$SIMD_BIN" ]; then echo "SIMD_BIN is not set. Make sure to run make install before"; exit 1; fi
echo "using $SIMD_BIN"
if [ -d "$($SIMD_BIN config home)" ]; then rm -r $($SIMD_BIN config home); fi
$SIMD_BIN config set client chain-id demo
$SIMD_BIN config set client keyring-backend test
$SIMD_BIN config set app api.enable true

echo $ALICE_MNEMONIC | $SIMD_BIN keys add alice --recover
echo $BOB_MNEMONIC | $SIMD_BIN keys add bob --recover
echo $CHARLIE_MNEMONIC | $SIMD_BIN keys add charlie --recover
echo $DAVID_MNEMONIC | $SIMD_BIN keys add david --recover
echo $UNKNOWN_MNOMONIC | $SIMD_BIN keys add unknown --recover

$SIMD_BIN init test --chain-id demo
$SIMD_BIN genesis add-genesis-account alice 10000000000000stake --keyring-backend test
$SIMD_BIN genesis add-genesis-account bob 5000000000stake --keyring-backend test
$SIMD_BIN genesis add-genesis-account charlie 5000000000stake --keyring-backend test
$SIMD_BIN genesis add-genesis-account david 5000000000stake --keyring-backend test
$SIMD_BIN genesis add-genesis-account unknown 5000000000stake --keyring-backend test

$SIMD_BIN genesis gentx alice 1000000000000stake --chain-id demo --keyring-backend test
$SIMD_BIN genesis collect-gentxs
$SIMD_BIN genesis validate-genesis