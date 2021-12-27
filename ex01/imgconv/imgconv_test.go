/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   imgconv_test.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/25 12:48:06 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/27 11:53:41 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package imgconv

import (
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		src   string
		dst   string
		rmsrc bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "test1",
			args:    args{src: "../testdata/pattern10_g.jpg", dst: "../testdata/pattern10_g.png", rmsrc: false},
			wantErr: nil,
		},
		{
			name:    "test2",
			args:    args{src: "../testdata/back_player.jpg", dst: "../testdata/back_player.png", rmsrc: false},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert(tt.args.src, tt.args.dst, tt.args.rmsrc); err != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v!!!!!!", err, tt.wantErr)
			} else {
				println("OK!!")
			}
		})
	}
}
