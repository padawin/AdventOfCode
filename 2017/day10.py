def get_ascii_codes(input_str):
    return [ord(c) for c in input_str] + [17, 31, 73, 47, 23]


def calculate_sparse_hash(numbers, lengths):
    nb_numbers = len(numbers)
    position = 0
    skip = 0
    for r in range(64):
        for length in lengths:
            # reverse order
            tmp = numbers + numbers
            start = position + length - 1
            end = (position - 1) if position is not 0 else None
            reversed_range = tmp[start:end:-1]
            for i, val in enumerate(reversed_range):
                numbers[(position + i) % nb_numbers] = val

            # update position
            position = (position + length + skip) % nb_numbers
            skip += 1

    return numbers


def calculate_dense_hash(h):
    return [h[  0] ^ h[  1] ^ h[  2] ^ h[  3] ^ h[  4] ^ h[  5] ^ h[  6] ^ h[  7] ^ h[  8] ^ h[  9] ^ h[ 10] ^ h[ 11] ^ h[ 12] ^ h[ 13] ^ h[ 14] ^ h[ 15],
            h[ 16] ^ h[ 17] ^ h[ 18] ^ h[ 19] ^ h[ 20] ^ h[ 21] ^ h[ 22] ^ h[ 23] ^ h[ 24] ^ h[ 25] ^ h[ 26] ^ h[ 27] ^ h[ 28] ^ h[ 29] ^ h[ 30] ^ h[ 31],
            h[ 32] ^ h[ 33] ^ h[ 34] ^ h[ 35] ^ h[ 36] ^ h[ 37] ^ h[ 38] ^ h[ 39] ^ h[ 40] ^ h[ 41] ^ h[ 42] ^ h[ 43] ^ h[ 44] ^ h[ 45] ^ h[ 46] ^ h[ 47],
            h[ 48] ^ h[ 49] ^ h[ 50] ^ h[ 51] ^ h[ 52] ^ h[ 53] ^ h[ 54] ^ h[ 55] ^ h[ 56] ^ h[ 57] ^ h[ 58] ^ h[ 59] ^ h[ 60] ^ h[ 61] ^ h[ 62] ^ h[ 63],
            h[ 64] ^ h[ 65] ^ h[ 66] ^ h[ 67] ^ h[ 68] ^ h[ 69] ^ h[ 70] ^ h[ 71] ^ h[ 72] ^ h[ 73] ^ h[ 74] ^ h[ 75] ^ h[ 76] ^ h[ 77] ^ h[ 78] ^ h[ 79],
            h[ 80] ^ h[ 81] ^ h[ 82] ^ h[ 83] ^ h[ 84] ^ h[ 85] ^ h[ 86] ^ h[ 87] ^ h[ 88] ^ h[ 89] ^ h[ 90] ^ h[ 91] ^ h[ 92] ^ h[ 93] ^ h[ 94] ^ h[ 95],
            h[ 96] ^ h[ 97] ^ h[ 98] ^ h[ 99] ^ h[100] ^ h[101] ^ h[102] ^ h[103] ^ h[104] ^ h[105] ^ h[106] ^ h[107] ^ h[108] ^ h[109] ^ h[110] ^ h[111],
            h[112] ^ h[113] ^ h[114] ^ h[115] ^ h[116] ^ h[117] ^ h[118] ^ h[119] ^ h[120] ^ h[121] ^ h[122] ^ h[123] ^ h[124] ^ h[125] ^ h[126] ^ h[127],
            h[128] ^ h[129] ^ h[130] ^ h[131] ^ h[132] ^ h[133] ^ h[134] ^ h[135] ^ h[136] ^ h[137] ^ h[138] ^ h[139] ^ h[140] ^ h[141] ^ h[142] ^ h[143],
            h[144] ^ h[145] ^ h[146] ^ h[147] ^ h[148] ^ h[149] ^ h[150] ^ h[151] ^ h[152] ^ h[153] ^ h[154] ^ h[155] ^ h[156] ^ h[157] ^ h[158] ^ h[159],
            h[160] ^ h[161] ^ h[162] ^ h[163] ^ h[164] ^ h[165] ^ h[166] ^ h[167] ^ h[168] ^ h[169] ^ h[170] ^ h[171] ^ h[172] ^ h[173] ^ h[174] ^ h[175],
            h[176] ^ h[177] ^ h[178] ^ h[179] ^ h[180] ^ h[181] ^ h[182] ^ h[183] ^ h[184] ^ h[185] ^ h[186] ^ h[187] ^ h[188] ^ h[189] ^ h[190] ^ h[191],
            h[192] ^ h[193] ^ h[194] ^ h[195] ^ h[196] ^ h[197] ^ h[198] ^ h[199] ^ h[200] ^ h[201] ^ h[202] ^ h[203] ^ h[204] ^ h[205] ^ h[206] ^ h[207],
            h[208] ^ h[209] ^ h[210] ^ h[211] ^ h[212] ^ h[213] ^ h[214] ^ h[215] ^ h[216] ^ h[217] ^ h[218] ^ h[219] ^ h[220] ^ h[221] ^ h[222] ^ h[223],
            h[224] ^ h[225] ^ h[226] ^ h[227] ^ h[228] ^ h[229] ^ h[230] ^ h[231] ^ h[232] ^ h[233] ^ h[234] ^ h[235] ^ h[236] ^ h[237] ^ h[238] ^ h[239],
            h[240] ^ h[241] ^ h[242] ^ h[243] ^ h[244] ^ h[245] ^ h[246] ^ h[247] ^ h[248] ^ h[249] ^ h[250] ^ h[251] ^ h[252] ^ h[253] ^ h[254] ^ h[255]]


def hash_list_to_str(hash):
    return ''.join("{:02x}".format(n) for n in hash)


def knot(s):
    ascii = get_ascii_codes(s)
    sparse = calculate_sparse_hash(list(range(256)), ascii)
    assert len(sparse) == 256
    dense = calculate_dense_hash(sparse)
    assert len(dense) == 16
    return hash_list_to_str(dense)


if __name__ == '__main__':
    input_vals = '225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110'
    assert knot("") == "a2582a3a0e66e6e86e3812dcb672a272"
    assert knot("AoC 2017") == "33efeb34ea91902bb2f59c9920caa6cd"
    assert knot("1,2,3") == "3efbe78a8d82f29979031a4aa0b16a9d"
    assert knot("1,2,4") == "63960835bcdc130f0b66d7ff4f6a5a8e"
    print(knot(input_vals))
