/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   imgconv_test.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/25 12:48:06 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/28 11:33:31 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package imgconv

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	t.Helper()
	type args struct {
		src   string
		dst   string
		rmsrc bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{src: "../testdata/pattern10_g.jpg", dst: "../testdata/pattern10_g.png", rmsrc: false},
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{src: "../testdata/back_player.jpg", dst: "../testdata/back_player.png", rmsrc: false},
			wantErr: false,
		},
		{
			name:    "test3",
			args:    args{src: "../testdata/back_player.j", dst: "../testdata/back_player.png", rmsrc: false},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := Convert(tt.args.src, tt.args.dst, tt.args.rmsrc); (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v!!!!!!", err, tt.wantErr)
			} else {
				fmt.Printf("\033[32mOK\033[0m\n")
			}
		})
	}
}
