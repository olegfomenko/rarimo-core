package cli

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudflare/bn256"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/rarimo-core/x/cbank/pkg"
	"math/big"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/rarimo/rarimo-core/x/cbank/types"
	"github.com/spf13/cobra"
)

func CmdGenerateKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-keypair",
		Short: "Generate BN256 keypair",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			params, err := queryClient.Params(context.TODO(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			prv, pub, err := pkg.PrivateKey(params.Params.G.MustToBN256G1())
			if err != nil {
				return err
			}

			cmd.Println("Private key:", hexutil.Encode(prv.Bytes()))
			cmd.Println("Public key", hexutil.Encode(pub.Marshal()))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGenerateCommitment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-keypair [v] [a1] [a2] [denom]",
		Short: "Generate commitment",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			arg_v := args[0]
			arg_a1 := args[1]
			arg_a2 := args[2]
			argDenom := args[3]

			params, err := queryClient.Params(context.TODO(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			a1 := new(big.Int).SetBytes(hexutil.MustDecode(arg_a1))
			a2 := new(big.Int).SetBytes(hexutil.MustDecode(arg_a2))

			v, ok := new(big.Int).SetString(arg_v, 10)
			if !ok {
				return errors.New("invalid v: should be in dec")
			}

			A1 := pkg.PublicKey(a1, params.Params.G.MustToBN256G1())
			A2 := pkg.PublicKey(a2, params.Params.G.MustToBN256G1())

			com := new(bn256.G1).Add(A1, pkg.PublicKey(v, params.Params.HVec[0].MustToBN256G1()))

			commitment := types.NewCommitment(com, A2, argDenom)
			cmd.Println("Commitment index:", commitment.Index())

			json, err := json.MarshalIndent(commitment, " ", "\t")
			if err != nil {
				return err
			}
			cmd.Println("Commitment entity:", string(json))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdPrepareReceiver() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare-receiver [A2] [v] [b1] [b2]",
		Short: "prepare 1-to-1 transfer for receiver size",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argA2 := args[0]
			arg_v := args[1]
			arg_b1 := args[2]
			arg_b2 := args[3]

			A2, err := new(types.Point).Decompress(argA2)
			if err != nil {
				return err
			}

			b1 := new(big.Int).SetBytes(hexutil.MustDecode(arg_b1))
			b2 := new(big.Int).SetBytes(hexutil.MustDecode(arg_b2))

			params, err := queryClient.Params(context.TODO(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			B2 := pkg.PublicKey(b2, params.Params.G.MustToBN256G1())

			v, ok := new(big.Int).SetString(arg_v, 10)
			if !ok {
				return errors.New("invalid v: should be in dec")
			}

			k1 := pkg.HashDH(A2.MustToBN256G1(), b1)
			k2 := pkg.HashDH(A2.MustToBN256G1(), b2)

			comOut := new(bn256.G1).Add(pkg.PublicKey(k1, params.Params.G.MustToBN256G1()), pkg.PublicKey(v, params.Params.HVec[0].MustToBN256G1()))
			addrOut := new(bn256.G1).Add(pkg.PublicKey(k2, params.Params.G.MustToBN256G1()), B2)

			// demon is not required for index
			commitment := types.NewCommitment(comOut, addrOut, "")

			cmd.Println("Commitment index:", commitment.Index())
			cmd.Println("Address", hexutil.Encode(addrOut.Marshal()))
			cmd.Println("Commitment private key", hexutil.Encode(k1.Bytes()))
			cmd.Println("Address private key", hexutil.Encode(new(big.Int).Add(k2, b2).Bytes()))
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
