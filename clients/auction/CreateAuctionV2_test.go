// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction

import (
	"bytes"
	"strconv"
	"testing"

	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
)

func TestEncodeDecode_CreateAuctionV2(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("CreateAuctionV2"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(CreateAuctionV2)
				fu.Fuzz(params)
				params.Args.Winners = &WinnerLimitCapped{
					Limit: 33,
				}
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				//
				got := new(CreateAuctionV2)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}