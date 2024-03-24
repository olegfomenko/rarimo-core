package cli

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"github.com/cloudflare/bn256"
	"github.com/rarimo/rarimo-core/x/cbank/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryGenerateParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-params",
		Short: "generate params",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, g, err := bn256.RandomG1(rand.Reader)
			if err != nil {
				return err
			}

			gvec := make([]*bn256.G1, 16)
			for i := range gvec {
				_, gvec[i], err = bn256.RandomG1(rand.Reader)
				if err != nil {
					return err
				}
			}

			hvec := make([]*bn256.G1, 32)
			for i := range hvec {
				_, hvec[i], err = bn256.RandomG1(rand.Reader)
				if err != nil {
					return err
				}
			}

			G := new(types.Point).MustFromBN256G1(g)
			GVec := new(types.PointSlice).MustFromBN256G1(gvec)
			HVec := new(types.PointSlice).MustFromBN256G1(hvec)

			params := types.Params{
				G:        *G,
				GVec:     (*GVec),
				HVec:     (*HVec)[:26],
				Nd:       16,
				Np:       15,
				GVecWNLA: []types.Point{},
				HVecWNLA: (*HVec)[26:],
			}

			json, err := json.MarshalIndent(params, " ", "\t")
			if err != nil {
				return err
			}
			cmd.Println(string(json))

			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
