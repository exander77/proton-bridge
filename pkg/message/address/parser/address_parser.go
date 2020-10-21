// Code generated from pkg/message/address/parser/Address.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Address

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 295,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	3, 2, 3, 2, 5, 2, 70, 10, 2, 3, 3, 7, 3, 73, 10, 3, 12, 3, 14, 3, 76, 11,
	3, 3, 3, 3, 3, 3, 3, 6, 3, 81, 10, 3, 13, 3, 14, 3, 82, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 5, 5, 90, 10, 5, 3, 6, 3, 6, 5, 6, 94, 10, 6, 3, 6, 7, 6,
	97, 10, 6, 12, 6, 14, 6, 100, 11, 6, 3, 6, 5, 6, 103, 10, 6, 3, 6, 3, 6,
	3, 7, 5, 7, 108, 10, 7, 3, 7, 6, 7, 111, 10, 7, 13, 7, 14, 7, 112, 3, 7,
	5, 7, 116, 10, 7, 3, 7, 5, 7, 119, 10, 7, 3, 8, 3, 8, 3, 9, 5, 9, 124,
	10, 9, 3, 9, 6, 9, 127, 10, 9, 13, 9, 14, 9, 128, 3, 9, 5, 9, 132, 10,
	9, 3, 10, 6, 10, 135, 10, 10, 13, 10, 14, 10, 136, 3, 10, 3, 10, 6, 10,
	141, 10, 10, 13, 10, 14, 10, 142, 7, 10, 145, 10, 10, 12, 10, 14, 10, 148,
	11, 10, 3, 11, 5, 11, 151, 10, 11, 3, 11, 3, 11, 5, 11, 155, 10, 11, 3,
	12, 3, 12, 3, 13, 3, 13, 5, 13, 161, 10, 13, 3, 14, 5, 14, 164, 10, 14,
	3, 14, 3, 14, 5, 14, 168, 10, 14, 3, 14, 7, 14, 171, 10, 14, 12, 14, 14,
	14, 174, 11, 14, 3, 14, 5, 14, 177, 10, 14, 3, 14, 3, 14, 5, 14, 181, 10,
	14, 3, 15, 3, 15, 5, 15, 185, 10, 15, 3, 16, 6, 16, 188, 10, 16, 13, 16,
	14, 16, 189, 3, 17, 5, 17, 193, 10, 17, 3, 17, 7, 17, 196, 10, 17, 12,
	17, 14, 17, 199, 11, 17, 3, 17, 7, 17, 202, 10, 17, 12, 17, 14, 17, 205,
	11, 17, 3, 18, 3, 18, 5, 18, 209, 10, 18, 3, 19, 3, 19, 5, 19, 213, 10,
	19, 3, 20, 5, 20, 216, 10, 20, 3, 20, 3, 20, 3, 21, 5, 21, 221, 10, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 227, 10, 21, 3, 22, 3, 22, 3, 22, 5,
	22, 232, 10, 22, 3, 22, 3, 22, 5, 22, 236, 10, 22, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 7, 24, 243, 10, 24, 12, 24, 14, 24, 246, 11, 24, 3, 25, 3,
	25, 3, 25, 7, 25, 251, 10, 25, 12, 25, 14, 25, 254, 11, 25, 3, 26, 3, 26,
	5, 26, 258, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 5, 28, 266,
	10, 28, 3, 29, 3, 29, 5, 29, 270, 10, 29, 3, 30, 5, 30, 273, 10, 30, 3,
	30, 3, 30, 5, 30, 277, 10, 30, 3, 30, 7, 30, 280, 10, 30, 12, 30, 14, 30,
	283, 11, 30, 3, 30, 3, 30, 5, 30, 287, 10, 30, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 2, 2, 34, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 2, 8, 5, 2, 7, 13, 16, 31, 33, 41, 11, 2, 7, 7, 9, 13,
	16, 17, 19, 19, 21, 22, 26, 26, 28, 28, 30, 30, 34, 41, 5, 2, 7, 7, 9,
	31, 33, 41, 4, 2, 7, 30, 34, 41, 3, 2, 5, 6, 3, 2, 7, 41, 2, 309, 2, 66,
	3, 2, 2, 2, 4, 74, 3, 2, 2, 2, 6, 84, 3, 2, 2, 2, 8, 89, 3, 2, 2, 2, 10,
	91, 3, 2, 2, 2, 12, 118, 3, 2, 2, 2, 14, 120, 3, 2, 2, 2, 16, 123, 3, 2,
	2, 2, 18, 134, 3, 2, 2, 2, 20, 150, 3, 2, 2, 2, 22, 156, 3, 2, 2, 2, 24,
	160, 3, 2, 2, 2, 26, 163, 3, 2, 2, 2, 28, 184, 3, 2, 2, 2, 30, 187, 3,
	2, 2, 2, 32, 197, 3, 2, 2, 2, 34, 208, 3, 2, 2, 2, 36, 212, 3, 2, 2, 2,
	38, 215, 3, 2, 2, 2, 40, 220, 3, 2, 2, 2, 42, 228, 3, 2, 2, 2, 44, 237,
	3, 2, 2, 2, 46, 239, 3, 2, 2, 2, 48, 247, 3, 2, 2, 2, 50, 257, 3, 2, 2,
	2, 52, 259, 3, 2, 2, 2, 54, 265, 3, 2, 2, 2, 56, 269, 3, 2, 2, 2, 58, 272,
	3, 2, 2, 2, 60, 288, 3, 2, 2, 2, 62, 290, 3, 2, 2, 2, 64, 292, 3, 2, 2,
	2, 66, 69, 7, 32, 2, 2, 67, 70, 5, 64, 33, 2, 68, 70, 5, 62, 32, 2, 69,
	67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 3, 3, 2, 2, 2, 71, 73, 5, 62, 32,
	2, 72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75,
	3, 2, 2, 2, 75, 77, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 78, 7, 3, 2, 2,
	78, 80, 3, 2, 2, 2, 79, 81, 5, 62, 32, 2, 80, 79, 3, 2, 2, 2, 81, 82, 3,
	2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 5, 3, 2, 2, 2, 84,
	85, 9, 2, 2, 2, 85, 7, 3, 2, 2, 2, 86, 90, 5, 6, 4, 2, 87, 90, 5, 2, 2,
	2, 88, 90, 5, 10, 6, 2, 89, 86, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 88,
	3, 2, 2, 2, 90, 9, 3, 2, 2, 2, 91, 98, 7, 14, 2, 2, 92, 94, 5, 4, 3, 2,
	93, 92, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 97, 5,
	8, 5, 2, 96, 93, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98,
	99, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 103, 5,
	4, 3, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 3, 2, 2,
	2, 104, 105, 7, 15, 2, 2, 105, 11, 3, 2, 2, 2, 106, 108, 5, 4, 3, 2, 107,
	106, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111,
	5, 10, 6, 2, 110, 107, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 110, 3, 2,
	2, 2, 112, 113, 3, 2, 2, 2, 113, 115, 3, 2, 2, 2, 114, 116, 5, 4, 3, 2,
	115, 114, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117,
	119, 5, 4, 3, 2, 118, 110, 3, 2, 2, 2, 118, 117, 3, 2, 2, 2, 119, 13, 3,
	2, 2, 2, 120, 121, 9, 3, 2, 2, 121, 15, 3, 2, 2, 2, 122, 124, 5, 12, 7,
	2, 123, 122, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126, 3, 2, 2, 2, 125,
	127, 5, 14, 8, 2, 126, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 126,
	3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 132, 5, 12,
	7, 2, 131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 17, 3, 2, 2, 2,
	133, 135, 5, 14, 8, 2, 134, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 146, 3, 2, 2, 2, 138, 140,
	7, 20, 2, 2, 139, 141, 5, 14, 8, 2, 140, 139, 3, 2, 2, 2, 141, 142, 3,
	2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2, 2,
	2, 144, 138, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 19, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 151, 5,
	12, 7, 2, 150, 149, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 152, 3, 2, 2,
	2, 152, 154, 5, 18, 10, 2, 153, 155, 5, 12, 7, 2, 154, 153, 3, 2, 2, 2,
	154, 155, 3, 2, 2, 2, 155, 21, 3, 2, 2, 2, 156, 157, 9, 4, 2, 2, 157, 23,
	3, 2, 2, 2, 158, 161, 5, 22, 12, 2, 159, 161, 5, 2, 2, 2, 160, 158, 3,
	2, 2, 2, 160, 159, 3, 2, 2, 2, 161, 25, 3, 2, 2, 2, 162, 164, 5, 12, 7,
	2, 163, 162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165,
	172, 7, 8, 2, 2, 166, 168, 5, 4, 3, 2, 167, 166, 3, 2, 2, 2, 167, 168,
	3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 171, 5, 24, 13, 2, 170, 167, 3,
	2, 2, 2, 171, 174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2,
	2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 175, 177, 5, 4, 3, 2, 176,
	175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180,
	7, 8, 2, 2, 179, 181, 5, 12, 7, 2, 180, 179, 3, 2, 2, 2, 180, 181, 3, 2,
	2, 2, 181, 27, 3, 2, 2, 2, 182, 185, 5, 16, 9, 2, 183, 185, 5, 26, 14,
	2, 184, 182, 3, 2, 2, 2, 184, 183, 3, 2, 2, 2, 185, 29, 3, 2, 2, 2, 186,
	188, 5, 28, 15, 2, 187, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 187,
	3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 31, 3, 2, 2, 2, 191, 193, 5, 4,
	3, 2, 192, 191, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2,
	194, 196, 5, 64, 33, 2, 195, 192, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197,
	195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 203, 3, 2, 2, 2, 199, 197,
	3, 2, 2, 2, 200, 202, 5, 62, 32, 2, 201, 200, 3, 2, 2, 2, 202, 205, 3,
	2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 33, 3, 2, 2,
	2, 205, 203, 3, 2, 2, 2, 206, 209, 5, 36, 19, 2, 207, 209, 5, 42, 22, 2,
	208, 206, 3, 2, 2, 2, 208, 207, 3, 2, 2, 2, 209, 35, 3, 2, 2, 2, 210, 213,
	5, 38, 20, 2, 211, 213, 5, 52, 27, 2, 212, 210, 3, 2, 2, 2, 212, 211, 3,
	2, 2, 2, 213, 37, 3, 2, 2, 2, 214, 216, 5, 44, 23, 2, 215, 214, 3, 2, 2,
	2, 215, 216, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 5, 40, 21, 2,
	218, 39, 3, 2, 2, 2, 219, 221, 5, 12, 7, 2, 220, 219, 3, 2, 2, 2, 220,
	221, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223, 7, 25, 2, 2, 223, 224,
	5, 52, 27, 2, 224, 226, 7, 27, 2, 2, 225, 227, 5, 12, 7, 2, 226, 225, 3,
	2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 41, 3, 2, 2, 2, 228, 229, 5, 44, 23,
	2, 229, 231, 7, 23, 2, 2, 230, 232, 5, 50, 26, 2, 231, 230, 3, 2, 2, 2,
	231, 232, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 235, 7, 24, 2, 2, 234,
	236, 5, 12, 7, 2, 235, 234, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 43,
	3, 2, 2, 2, 237, 238, 5, 30, 16, 2, 238, 45, 3, 2, 2, 2, 239, 244, 5, 36,
	19, 2, 240, 241, 7, 18, 2, 2, 241, 243, 5, 36, 19, 2, 242, 240, 3, 2, 2,
	2, 243, 246, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245,
	47, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 247, 252, 5, 34, 18, 2, 248, 249,
	7, 18, 2, 2, 249, 251, 5, 34, 18, 2, 250, 248, 3, 2, 2, 2, 251, 254, 3,
	2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 49, 3, 2, 2,
	2, 254, 252, 3, 2, 2, 2, 255, 258, 5, 46, 24, 2, 256, 258, 5, 12, 7, 2,
	257, 255, 3, 2, 2, 2, 257, 256, 3, 2, 2, 2, 258, 51, 3, 2, 2, 2, 259, 260,
	5, 54, 28, 2, 260, 261, 7, 29, 2, 2, 261, 262, 5, 56, 29, 2, 262, 53, 3,
	2, 2, 2, 263, 266, 5, 20, 11, 2, 264, 266, 5, 26, 14, 2, 265, 263, 3, 2,
	2, 2, 265, 264, 3, 2, 2, 2, 266, 55, 3, 2, 2, 2, 267, 270, 5, 20, 11, 2,
	268, 270, 5, 58, 30, 2, 269, 267, 3, 2, 2, 2, 269, 268, 3, 2, 2, 2, 270,
	57, 3, 2, 2, 2, 271, 273, 5, 12, 7, 2, 272, 271, 3, 2, 2, 2, 272, 273,
	3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 281, 7, 31, 2, 2, 275, 277, 5, 4,
	3, 2, 276, 275, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2,
	278, 280, 5, 60, 31, 2, 279, 276, 3, 2, 2, 2, 280, 283, 3, 2, 2, 2, 281,
	279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 284, 3, 2, 2, 2, 283, 281,
	3, 2, 2, 2, 284, 286, 7, 33, 2, 2, 285, 287, 5, 12, 7, 2, 286, 285, 3,
	2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 59, 3, 2, 2, 2, 288, 289, 9, 5, 2,
	2, 289, 61, 3, 2, 2, 2, 290, 291, 9, 6, 2, 2, 291, 63, 3, 2, 2, 2, 292,
	293, 9, 7, 2, 2, 293, 65, 3, 2, 2, 2, 48, 69, 74, 82, 89, 93, 98, 102,
	107, 112, 115, 118, 123, 128, 131, 136, 142, 146, 150, 154, 160, 163, 167,
	172, 176, 180, 184, 189, 192, 197, 203, 208, 212, 215, 220, 226, 231, 235,
	244, 252, 257, 265, 269, 272, 276, 281, 286,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\r\n'", "'\u007F'", "'\t'", "' '", "'!'", "'\"'", "'#'", "'$'", "'%'",
	"'&'", "'''", "'('", "')'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "",
	"':'", "';'", "'<'", "'='", "'>'", "'?'", "'@'", "", "'['", "'\\'", "']'",
	"'^'", "'_'", "'`'", "", "'{'", "'|'", "'}'", "'~'",
}
var symbolicNames = []string{
	"", "CRLF", "DEL", "HTAB", "SP", "Exclamation", "DQuote", "Hash", "Dollar",
	"Percent", "Ampersand", "SQuote", "LParens", "RParens", "Asterisk", "Plus",
	"Comma", "Minus", "Period", "Slash", "Digit", "Colon", "Semicolon", "Less",
	"Equal", "Greater", "Question", "At", "AlphaUpper", "LBracket", "Backslash",
	"RBracket", "Caret", "Underscore", "Backtick", "AlphaLower", "LCurly",
	"Pipe", "RCurly", "Tilde",
}

var ruleNames = []string{
	"quotedPair", "fws", "ctext", "ccontent", "comment", "cfws", "atext", "atom",
	"dotAtomText", "dotAtom", "qtext", "qcontent", "quotedString", "word",
	"phrase", "unstructured", "address", "mailbox", "nameAddr", "angleAddr",
	"group", "displayName", "mailboxList", "addressList", "groupList", "addrSpec",
	"localPart", "domain", "domainLiteral", "dtext", "wsp", "vchar",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Address struct {
	*antlr.BaseParser
}

func NewAddress(input antlr.TokenStream) *Address {
	this := new(Address)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Address.g4"

	return this
}

// Address tokens.
const (
	AddressEOF         = antlr.TokenEOF
	AddressCRLF        = 1
	AddressDEL         = 2
	AddressHTAB        = 3
	AddressSP          = 4
	AddressExclamation = 5
	AddressDQuote      = 6
	AddressHash        = 7
	AddressDollar      = 8
	AddressPercent     = 9
	AddressAmpersand   = 10
	AddressSQuote      = 11
	AddressLParens     = 12
	AddressRParens     = 13
	AddressAsterisk    = 14
	AddressPlus        = 15
	AddressComma       = 16
	AddressMinus       = 17
	AddressPeriod      = 18
	AddressSlash       = 19
	AddressDigit       = 20
	AddressColon       = 21
	AddressSemicolon   = 22
	AddressLess        = 23
	AddressEqual       = 24
	AddressGreater     = 25
	AddressQuestion    = 26
	AddressAt          = 27
	AddressAlphaUpper  = 28
	AddressLBracket    = 29
	AddressBackslash   = 30
	AddressRBracket    = 31
	AddressCaret       = 32
	AddressUnderscore  = 33
	AddressBacktick    = 34
	AddressAlphaLower  = 35
	AddressLCurly      = 36
	AddressPipe        = 37
	AddressRCurly      = 38
	AddressTilde       = 39
)

// Address rules.
const (
	AddressRULE_quotedPair    = 0
	AddressRULE_fws           = 1
	AddressRULE_ctext         = 2
	AddressRULE_ccontent      = 3
	AddressRULE_comment       = 4
	AddressRULE_cfws          = 5
	AddressRULE_atext         = 6
	AddressRULE_atom          = 7
	AddressRULE_dotAtomText   = 8
	AddressRULE_dotAtom       = 9
	AddressRULE_qtext         = 10
	AddressRULE_qcontent      = 11
	AddressRULE_quotedString  = 12
	AddressRULE_word          = 13
	AddressRULE_phrase        = 14
	AddressRULE_unstructured  = 15
	AddressRULE_address       = 16
	AddressRULE_mailbox       = 17
	AddressRULE_nameAddr      = 18
	AddressRULE_angleAddr     = 19
	AddressRULE_group         = 20
	AddressRULE_displayName   = 21
	AddressRULE_mailboxList   = 22
	AddressRULE_addressList   = 23
	AddressRULE_groupList     = 24
	AddressRULE_addrSpec      = 25
	AddressRULE_localPart     = 26
	AddressRULE_domain        = 27
	AddressRULE_domainLiteral = 28
	AddressRULE_dtext         = 29
	AddressRULE_wsp           = 30
	AddressRULE_vchar         = 31
)

// IQuotedPairContext is an interface to support dynamic dispatch.
type IQuotedPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedPairContext differentiates from other interfaces.
	IsQuotedPairContext()
}

type QuotedPairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedPairContext() *QuotedPairContext {
	var p = new(QuotedPairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_quotedPair
	return p
}

func (*QuotedPairContext) IsQuotedPairContext() {}

func NewQuotedPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedPairContext {
	var p = new(QuotedPairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_quotedPair

	return p
}

func (s *QuotedPairContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedPairContext) Backslash() antlr.TerminalNode {
	return s.GetToken(AddressBackslash, 0)
}

func (s *QuotedPairContext) Vchar() IVcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVcharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVcharContext)
}

func (s *QuotedPairContext) Wsp() IWspContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWspContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWspContext)
}

func (s *QuotedPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterQuotedPair(s)
	}
}

func (s *QuotedPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitQuotedPair(s)
	}
}

func (p *Address) QuotedPair() (localctx IQuotedPairContext) {
	localctx = NewQuotedPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AddressRULE_quotedPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(AddressBackslash)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressExclamation, AddressDQuote, AddressHash, AddressDollar, AddressPercent, AddressAmpersand, AddressSQuote, AddressLParens, AddressRParens, AddressAsterisk, AddressPlus, AddressComma, AddressMinus, AddressPeriod, AddressSlash, AddressDigit, AddressColon, AddressSemicolon, AddressLess, AddressEqual, AddressGreater, AddressQuestion, AddressAt, AddressAlphaUpper, AddressLBracket, AddressBackslash, AddressRBracket, AddressCaret, AddressUnderscore, AddressBacktick, AddressAlphaLower, AddressLCurly, AddressPipe, AddressRCurly, AddressTilde:
		{
			p.SetState(65)
			p.Vchar()
		}

	case AddressHTAB, AddressSP:
		{
			p.SetState(66)
			p.Wsp()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFwsContext is an interface to support dynamic dispatch.
type IFwsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFwsContext differentiates from other interfaces.
	IsFwsContext()
}

type FwsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFwsContext() *FwsContext {
	var p = new(FwsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_fws
	return p
}

func (*FwsContext) IsFwsContext() {}

func NewFwsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FwsContext {
	var p = new(FwsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_fws

	return p
}

func (s *FwsContext) GetParser() antlr.Parser { return s.parser }

func (s *FwsContext) CRLF() antlr.TerminalNode {
	return s.GetToken(AddressCRLF, 0)
}

func (s *FwsContext) AllWsp() []IWspContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWspContext)(nil)).Elem())
	var tst = make([]IWspContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWspContext)
		}
	}

	return tst
}

func (s *FwsContext) Wsp(i int) IWspContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWspContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWspContext)
}

func (s *FwsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FwsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FwsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterFws(s)
	}
}

func (s *FwsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitFws(s)
	}
}

func (p *Address) Fws() (localctx IFwsContext) {
	localctx = NewFwsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AddressRULE_fws)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressHTAB || _la == AddressSP {
		{
			p.SetState(69)
			p.Wsp()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(AddressCRLF)
	}

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(77)
				p.Wsp()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ICtextContext is an interface to support dynamic dispatch.
type ICtextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCtextContext differentiates from other interfaces.
	IsCtextContext()
}

type CtextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtextContext() *CtextContext {
	var p = new(CtextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_ctext
	return p
}

func (*CtextContext) IsCtextContext() {}

func NewCtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CtextContext {
	var p = new(CtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_ctext

	return p
}

func (s *CtextContext) GetParser() antlr.Parser { return s.parser }

func (s *CtextContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressExclamation, 0)
}

func (s *CtextContext) DQuote() antlr.TerminalNode {
	return s.GetToken(AddressDQuote, 0)
}

func (s *CtextContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressHash, 0)
}

func (s *CtextContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressDollar, 0)
}

func (s *CtextContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressPercent, 0)
}

func (s *CtextContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressAmpersand, 0)
}

func (s *CtextContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressSQuote, 0)
}

func (s *CtextContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressAsterisk, 0)
}

func (s *CtextContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressPlus, 0)
}

func (s *CtextContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressComma, 0)
}

func (s *CtextContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressMinus, 0)
}

func (s *CtextContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressPeriod, 0)
}

func (s *CtextContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressSlash, 0)
}

func (s *CtextContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressDigit, 0)
}

func (s *CtextContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressColon, 0)
}

func (s *CtextContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressSemicolon, 0)
}

func (s *CtextContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressLess, 0)
}

func (s *CtextContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressEqual, 0)
}

func (s *CtextContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressGreater, 0)
}

func (s *CtextContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressQuestion, 0)
}

func (s *CtextContext) At() antlr.TerminalNode {
	return s.GetToken(AddressAt, 0)
}

func (s *CtextContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressAlphaUpper, 0)
}

func (s *CtextContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressLBracket, 0)
}

func (s *CtextContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressRBracket, 0)
}

func (s *CtextContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressCaret, 0)
}

func (s *CtextContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressUnderscore, 0)
}

func (s *CtextContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressBacktick, 0)
}

func (s *CtextContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressAlphaLower, 0)
}

func (s *CtextContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressLCurly, 0)
}

func (s *CtextContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressPipe, 0)
}

func (s *CtextContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressRCurly, 0)
}

func (s *CtextContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressTilde, 0)
}

func (s *CtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterCtext(s)
	}
}

func (s *CtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitCtext(s)
	}
}

func (p *Address) Ctext() (localctx ICtextContext) {
	localctx = NewCtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AddressRULE_ctext)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressExclamation)|(1<<AddressDQuote)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressComma)|(1<<AddressMinus)|(1<<AddressPeriod)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressColon)|(1<<AddressSemicolon)|(1<<AddressLess)|(1<<AddressEqual)|(1<<AddressGreater)|(1<<AddressQuestion)|(1<<AddressAt)|(1<<AddressAlphaUpper)|(1<<AddressLBracket)|(1<<AddressRBracket))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICcontentContext is an interface to support dynamic dispatch.
type ICcontentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCcontentContext differentiates from other interfaces.
	IsCcontentContext()
}

type CcontentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCcontentContext() *CcontentContext {
	var p = new(CcontentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_ccontent
	return p
}

func (*CcontentContext) IsCcontentContext() {}

func NewCcontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CcontentContext {
	var p = new(CcontentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_ccontent

	return p
}

func (s *CcontentContext) GetParser() antlr.Parser { return s.parser }

func (s *CcontentContext) Ctext() ICtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICtextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICtextContext)
}

func (s *CcontentContext) QuotedPair() IQuotedPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedPairContext)
}

func (s *CcontentContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *CcontentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CcontentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CcontentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterCcontent(s)
	}
}

func (s *CcontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitCcontent(s)
	}
}

func (p *Address) Ccontent() (localctx ICcontentContext) {
	localctx = NewCcontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AddressRULE_ccontent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressExclamation, AddressDQuote, AddressHash, AddressDollar, AddressPercent, AddressAmpersand, AddressSQuote, AddressAsterisk, AddressPlus, AddressComma, AddressMinus, AddressPeriod, AddressSlash, AddressDigit, AddressColon, AddressSemicolon, AddressLess, AddressEqual, AddressGreater, AddressQuestion, AddressAt, AddressAlphaUpper, AddressLBracket, AddressRBracket, AddressCaret, AddressUnderscore, AddressBacktick, AddressAlphaLower, AddressLCurly, AddressPipe, AddressRCurly, AddressTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Ctext()
		}

	case AddressBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.QuotedPair()
		}

	case AddressLParens:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressLParens, 0)
}

func (s *CommentContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressRParens, 0)
}

func (s *CommentContext) AllCcontent() []ICcontentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICcontentContext)(nil)).Elem())
	var tst = make([]ICcontentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICcontentContext)
		}
	}

	return tst
}

func (s *CommentContext) Ccontent(i int) ICcontentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICcontentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICcontentContext)
}

func (s *CommentContext) AllFws() []IFwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFwsContext)(nil)).Elem())
	var tst = make([]IFwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFwsContext)
		}
	}

	return tst
}

func (s *CommentContext) Fws(i int) IFwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFwsContext)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *Address) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AddressRULE_comment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(AddressLParens)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP))) != 0 {
				{
					p.SetState(90)
					p.Fws()
				}

			}
			{
				p.SetState(93)
				p.Ccontent()
			}

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP))) != 0 {
		{
			p.SetState(99)
			p.Fws()
		}

	}
	{
		p.SetState(102)
		p.Match(AddressRParens)
	}

	return localctx
}

// ICfwsContext is an interface to support dynamic dispatch.
type ICfwsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCfwsContext differentiates from other interfaces.
	IsCfwsContext()
}

type CfwsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCfwsContext() *CfwsContext {
	var p = new(CfwsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_cfws
	return p
}

func (*CfwsContext) IsCfwsContext() {}

func NewCfwsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CfwsContext {
	var p = new(CfwsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_cfws

	return p
}

func (s *CfwsContext) GetParser() antlr.Parser { return s.parser }

func (s *CfwsContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *CfwsContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *CfwsContext) AllFws() []IFwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFwsContext)(nil)).Elem())
	var tst = make([]IFwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFwsContext)
		}
	}

	return tst
}

func (s *CfwsContext) Fws(i int) IFwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFwsContext)
}

func (s *CfwsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CfwsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CfwsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterCfws(s)
	}
}

func (s *CfwsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitCfws(s)
	}
}

func (p *Address) Cfws() (localctx ICfwsContext) {
	localctx = NewCfwsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AddressRULE_cfws)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(105)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP))) != 0 {
					{
						p.SetState(104)
						p.Fws()
					}

				}
				{
					p.SetState(107)
					p.Comment()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(112)
				p.Fws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Fws()
		}

	}

	return localctx
}

// IAtextContext is an interface to support dynamic dispatch.
type IAtextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtextContext differentiates from other interfaces.
	IsAtextContext()
}

type AtextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtextContext() *AtextContext {
	var p = new(AtextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_atext
	return p
}

func (*AtextContext) IsAtextContext() {}

func NewAtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtextContext {
	var p = new(AtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_atext

	return p
}

func (s *AtextContext) GetParser() antlr.Parser { return s.parser }

func (s *AtextContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressAlphaUpper, 0)
}

func (s *AtextContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressAlphaLower, 0)
}

func (s *AtextContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressDigit, 0)
}

func (s *AtextContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressExclamation, 0)
}

func (s *AtextContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressHash, 0)
}

func (s *AtextContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressDollar, 0)
}

func (s *AtextContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressPercent, 0)
}

func (s *AtextContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressAmpersand, 0)
}

func (s *AtextContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressSQuote, 0)
}

func (s *AtextContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressAsterisk, 0)
}

func (s *AtextContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressPlus, 0)
}

func (s *AtextContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressMinus, 0)
}

func (s *AtextContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressSlash, 0)
}

func (s *AtextContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressEqual, 0)
}

func (s *AtextContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressQuestion, 0)
}

func (s *AtextContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressCaret, 0)
}

func (s *AtextContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressUnderscore, 0)
}

func (s *AtextContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressBacktick, 0)
}

func (s *AtextContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressLCurly, 0)
}

func (s *AtextContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressPipe, 0)
}

func (s *AtextContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressRCurly, 0)
}

func (s *AtextContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressTilde, 0)
}

func (s *AtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterAtext(s)
	}
}

func (s *AtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitAtext(s)
	}
}

func (p *Address) Atext() (localctx IAtextContext) {
	localctx = NewAtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AddressRULE_atext)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressExclamation)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressMinus)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressEqual)|(1<<AddressQuestion)|(1<<AddressAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *AtomContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *AtomContext) AllAtext() []IAtextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtextContext)(nil)).Elem())
	var tst = make([]IAtextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtextContext)
		}
	}

	return tst
}

func (s *AtomContext) Atext(i int) IAtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtextContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *Address) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AddressRULE_atom)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(120)
			p.Cfws()
		}

	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(123)
				p.Atext()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(128)
			p.Cfws()
		}

	}

	return localctx
}

// IDotAtomTextContext is an interface to support dynamic dispatch.
type IDotAtomTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDotAtomTextContext differentiates from other interfaces.
	IsDotAtomTextContext()
}

type DotAtomTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotAtomTextContext() *DotAtomTextContext {
	var p = new(DotAtomTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_dotAtomText
	return p
}

func (*DotAtomTextContext) IsDotAtomTextContext() {}

func NewDotAtomTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotAtomTextContext {
	var p = new(DotAtomTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_dotAtomText

	return p
}

func (s *DotAtomTextContext) GetParser() antlr.Parser { return s.parser }

func (s *DotAtomTextContext) AllAtext() []IAtextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtextContext)(nil)).Elem())
	var tst = make([]IAtextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtextContext)
		}
	}

	return tst
}

func (s *DotAtomTextContext) Atext(i int) IAtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtextContext)
}

func (s *DotAtomTextContext) AllPeriod() []antlr.TerminalNode {
	return s.GetTokens(AddressPeriod)
}

func (s *DotAtomTextContext) Period(i int) antlr.TerminalNode {
	return s.GetToken(AddressPeriod, i)
}

func (s *DotAtomTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotAtomTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DotAtomTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterDotAtomText(s)
	}
}

func (s *DotAtomTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitDotAtomText(s)
	}
}

func (p *Address) DotAtomText() (localctx IDotAtomTextContext) {
	localctx = NewDotAtomTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AddressRULE_dotAtomText)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressExclamation)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressMinus)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressEqual)|(1<<AddressQuestion)|(1<<AddressAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0) {
		{
			p.SetState(131)
			p.Atext()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressPeriod {
		{
			p.SetState(136)
			p.Match(AddressPeriod)
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressExclamation)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressMinus)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressEqual)|(1<<AddressQuestion)|(1<<AddressAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0) {
			{
				p.SetState(137)
				p.Atext()
			}

			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDotAtomContext is an interface to support dynamic dispatch.
type IDotAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDotAtomContext differentiates from other interfaces.
	IsDotAtomContext()
}

type DotAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotAtomContext() *DotAtomContext {
	var p = new(DotAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_dotAtom
	return p
}

func (*DotAtomContext) IsDotAtomContext() {}

func NewDotAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotAtomContext {
	var p = new(DotAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_dotAtom

	return p
}

func (s *DotAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *DotAtomContext) DotAtomText() IDotAtomTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotAtomTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDotAtomTextContext)
}

func (s *DotAtomContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *DotAtomContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *DotAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DotAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterDotAtom(s)
	}
}

func (s *DotAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitDotAtom(s)
	}
}

func (p *Address) DotAtom() (localctx IDotAtomContext) {
	localctx = NewDotAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AddressRULE_dotAtom)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(147)
			p.Cfws()
		}

	}
	{
		p.SetState(150)
		p.DotAtomText()
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(151)
			p.Cfws()
		}

	}

	return localctx
}

// IQtextContext is an interface to support dynamic dispatch.
type IQtextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQtextContext differentiates from other interfaces.
	IsQtextContext()
}

type QtextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQtextContext() *QtextContext {
	var p = new(QtextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_qtext
	return p
}

func (*QtextContext) IsQtextContext() {}

func NewQtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QtextContext {
	var p = new(QtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_qtext

	return p
}

func (s *QtextContext) GetParser() antlr.Parser { return s.parser }

func (s *QtextContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressExclamation, 0)
}

func (s *QtextContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressHash, 0)
}

func (s *QtextContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressDollar, 0)
}

func (s *QtextContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressPercent, 0)
}

func (s *QtextContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressAmpersand, 0)
}

func (s *QtextContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressSQuote, 0)
}

func (s *QtextContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressLParens, 0)
}

func (s *QtextContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressRParens, 0)
}

func (s *QtextContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressAsterisk, 0)
}

func (s *QtextContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressPlus, 0)
}

func (s *QtextContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressComma, 0)
}

func (s *QtextContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressMinus, 0)
}

func (s *QtextContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressPeriod, 0)
}

func (s *QtextContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressSlash, 0)
}

func (s *QtextContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressDigit, 0)
}

func (s *QtextContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressColon, 0)
}

func (s *QtextContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressSemicolon, 0)
}

func (s *QtextContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressLess, 0)
}

func (s *QtextContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressEqual, 0)
}

func (s *QtextContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressGreater, 0)
}

func (s *QtextContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressQuestion, 0)
}

func (s *QtextContext) At() antlr.TerminalNode {
	return s.GetToken(AddressAt, 0)
}

func (s *QtextContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressAlphaUpper, 0)
}

func (s *QtextContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressLBracket, 0)
}

func (s *QtextContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressRBracket, 0)
}

func (s *QtextContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressCaret, 0)
}

func (s *QtextContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressUnderscore, 0)
}

func (s *QtextContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressBacktick, 0)
}

func (s *QtextContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressAlphaLower, 0)
}

func (s *QtextContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressLCurly, 0)
}

func (s *QtextContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressPipe, 0)
}

func (s *QtextContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressRCurly, 0)
}

func (s *QtextContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressTilde, 0)
}

func (s *QtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterQtext(s)
	}
}

func (s *QtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitQtext(s)
	}
}

func (p *Address) Qtext() (localctx IQtextContext) {
	localctx = NewQtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AddressRULE_qtext)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressExclamation)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressLParens)|(1<<AddressRParens)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressComma)|(1<<AddressMinus)|(1<<AddressPeriod)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressColon)|(1<<AddressSemicolon)|(1<<AddressLess)|(1<<AddressEqual)|(1<<AddressGreater)|(1<<AddressQuestion)|(1<<AddressAt)|(1<<AddressAlphaUpper)|(1<<AddressLBracket)|(1<<AddressRBracket))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IQcontentContext is an interface to support dynamic dispatch.
type IQcontentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQcontentContext differentiates from other interfaces.
	IsQcontentContext()
}

type QcontentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQcontentContext() *QcontentContext {
	var p = new(QcontentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_qcontent
	return p
}

func (*QcontentContext) IsQcontentContext() {}

func NewQcontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QcontentContext {
	var p = new(QcontentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_qcontent

	return p
}

func (s *QcontentContext) GetParser() antlr.Parser { return s.parser }

func (s *QcontentContext) Qtext() IQtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQtextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQtextContext)
}

func (s *QcontentContext) QuotedPair() IQuotedPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedPairContext)
}

func (s *QcontentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QcontentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QcontentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterQcontent(s)
	}
}

func (s *QcontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitQcontent(s)
	}
}

func (p *Address) Qcontent() (localctx IQcontentContext) {
	localctx = NewQcontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AddressRULE_qcontent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressExclamation, AddressHash, AddressDollar, AddressPercent, AddressAmpersand, AddressSQuote, AddressLParens, AddressRParens, AddressAsterisk, AddressPlus, AddressComma, AddressMinus, AddressPeriod, AddressSlash, AddressDigit, AddressColon, AddressSemicolon, AddressLess, AddressEqual, AddressGreater, AddressQuestion, AddressAt, AddressAlphaUpper, AddressLBracket, AddressRBracket, AddressCaret, AddressUnderscore, AddressBacktick, AddressAlphaLower, AddressLCurly, AddressPipe, AddressRCurly, AddressTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Qtext()
		}

	case AddressBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.QuotedPair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuotedStringContext is an interface to support dynamic dispatch.
type IQuotedStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedStringContext differentiates from other interfaces.
	IsQuotedStringContext()
}

type QuotedStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedStringContext() *QuotedStringContext {
	var p = new(QuotedStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_quotedString
	return p
}

func (*QuotedStringContext) IsQuotedStringContext() {}

func NewQuotedStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedStringContext {
	var p = new(QuotedStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_quotedString

	return p
}

func (s *QuotedStringContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedStringContext) AllDQuote() []antlr.TerminalNode {
	return s.GetTokens(AddressDQuote)
}

func (s *QuotedStringContext) DQuote(i int) antlr.TerminalNode {
	return s.GetToken(AddressDQuote, i)
}

func (s *QuotedStringContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *QuotedStringContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *QuotedStringContext) AllQcontent() []IQcontentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQcontentContext)(nil)).Elem())
	var tst = make([]IQcontentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQcontentContext)
		}
	}

	return tst
}

func (s *QuotedStringContext) Qcontent(i int) IQcontentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQcontentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQcontentContext)
}

func (s *QuotedStringContext) AllFws() []IFwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFwsContext)(nil)).Elem())
	var tst = make([]IFwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFwsContext)
		}
	}

	return tst
}

func (s *QuotedStringContext) Fws(i int) IFwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFwsContext)
}

func (s *QuotedStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterQuotedString(s)
	}
}

func (s *QuotedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitQuotedString(s)
	}
}

func (p *Address) QuotedString() (localctx IQuotedStringContext) {
	localctx = NewQuotedStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AddressRULE_quotedString)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(160)
			p.Cfws()
		}

	}
	{
		p.SetState(163)
		p.Match(AddressDQuote)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP))) != 0 {
				{
					p.SetState(164)
					p.Fws()
				}

			}
			{
				p.SetState(167)
				p.Qcontent()
			}

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP))) != 0 {
		{
			p.SetState(173)
			p.Fws()
		}

	}
	{
		p.SetState(176)
		p.Match(AddressDQuote)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(177)
			p.Cfws()
		}

	}

	return localctx
}

// IWordContext is an interface to support dynamic dispatch.
type IWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordContext differentiates from other interfaces.
	IsWordContext()
}

type WordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordContext() *WordContext {
	var p = new(WordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_word
	return p
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_word

	return p
}

func (s *WordContext) GetParser() antlr.Parser { return s.parser }

func (s *WordContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *WordContext) QuotedString() IQuotedStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedStringContext)
}

func (s *WordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterWord(s)
	}
}

func (s *WordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitWord(s)
	}
}

func (p *Address) Word() (localctx IWordContext) {
	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AddressRULE_word)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Atom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.QuotedString()
		}

	}

	return localctx
}

// IPhraseContext is an interface to support dynamic dispatch.
type IPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhraseContext differentiates from other interfaces.
	IsPhraseContext()
}

type PhraseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhraseContext() *PhraseContext {
	var p = new(PhraseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_phrase
	return p
}

func (*PhraseContext) IsPhraseContext() {}

func NewPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhraseContext {
	var p = new(PhraseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_phrase

	return p
}

func (s *PhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *PhraseContext) AllWord() []IWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordContext)(nil)).Elem())
	var tst = make([]IWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordContext)
		}
	}

	return tst
}

func (s *PhraseContext) Word(i int) IWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *PhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterPhrase(s)
	}
}

func (s *PhraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitPhrase(s)
	}
}

func (p *Address) Phrase() (localctx IPhraseContext) {
	localctx = NewPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AddressRULE_phrase)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(184)
				p.Word()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// IUnstructuredContext is an interface to support dynamic dispatch.
type IUnstructuredContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnstructuredContext differentiates from other interfaces.
	IsUnstructuredContext()
}

type UnstructuredContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnstructuredContext() *UnstructuredContext {
	var p = new(UnstructuredContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_unstructured
	return p
}

func (*UnstructuredContext) IsUnstructuredContext() {}

func NewUnstructuredContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnstructuredContext {
	var p = new(UnstructuredContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_unstructured

	return p
}

func (s *UnstructuredContext) GetParser() antlr.Parser { return s.parser }

func (s *UnstructuredContext) AllVchar() []IVcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVcharContext)(nil)).Elem())
	var tst = make([]IVcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVcharContext)
		}
	}

	return tst
}

func (s *UnstructuredContext) Vchar(i int) IVcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVcharContext)
}

func (s *UnstructuredContext) AllWsp() []IWspContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWspContext)(nil)).Elem())
	var tst = make([]IWspContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWspContext)
		}
	}

	return tst
}

func (s *UnstructuredContext) Wsp(i int) IWspContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWspContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWspContext)
}

func (s *UnstructuredContext) AllFws() []IFwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFwsContext)(nil)).Elem())
	var tst = make([]IFwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFwsContext)
		}
	}

	return tst
}

func (s *UnstructuredContext) Fws(i int) IFwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFwsContext)
}

func (s *UnstructuredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnstructuredContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnstructuredContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterUnstructured(s)
	}
}

func (s *UnstructuredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitUnstructured(s)
	}
}

func (p *Address) Unstructured() (localctx IUnstructuredContext) {
	localctx = NewUnstructuredContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AddressRULE_unstructured)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP))) != 0 {
				{
					p.SetState(189)
					p.Fws()
				}

			}
			{
				p.SetState(192)
				p.Vchar()
			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressHTAB || _la == AddressSP {
		{
			p.SetState(198)
			p.Wsp()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAddressContext is an interface to support dynamic dispatch.
type IAddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddressContext differentiates from other interfaces.
	IsAddressContext()
}

type AddressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddressContext() *AddressContext {
	var p = new(AddressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_address
	return p
}

func (*AddressContext) IsAddressContext() {}

func NewAddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddressContext {
	var p = new(AddressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_address

	return p
}

func (s *AddressContext) GetParser() antlr.Parser { return s.parser }

func (s *AddressContext) Mailbox() IMailboxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMailboxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMailboxContext)
}

func (s *AddressContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *AddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterAddress(s)
	}
}

func (s *AddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitAddress(s)
	}
}

func (p *Address) Address() (localctx IAddressContext) {
	localctx = NewAddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AddressRULE_address)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.Mailbox()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.Group()
		}

	}

	return localctx
}

// IMailboxContext is an interface to support dynamic dispatch.
type IMailboxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMailboxContext differentiates from other interfaces.
	IsMailboxContext()
}

type MailboxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMailboxContext() *MailboxContext {
	var p = new(MailboxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_mailbox
	return p
}

func (*MailboxContext) IsMailboxContext() {}

func NewMailboxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MailboxContext {
	var p = new(MailboxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_mailbox

	return p
}

func (s *MailboxContext) GetParser() antlr.Parser { return s.parser }

func (s *MailboxContext) NameAddr() INameAddrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAddrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameAddrContext)
}

func (s *MailboxContext) AddrSpec() IAddrSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddrSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddrSpecContext)
}

func (s *MailboxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MailboxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MailboxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterMailbox(s)
	}
}

func (s *MailboxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitMailbox(s)
	}
}

func (p *Address) Mailbox() (localctx IMailboxContext) {
	localctx = NewMailboxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AddressRULE_mailbox)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.NameAddr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.AddrSpec()
		}

	}

	return localctx
}

// INameAddrContext is an interface to support dynamic dispatch.
type INameAddrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameAddrContext differentiates from other interfaces.
	IsNameAddrContext()
}

type NameAddrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameAddrContext() *NameAddrContext {
	var p = new(NameAddrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_nameAddr
	return p
}

func (*NameAddrContext) IsNameAddrContext() {}

func NewNameAddrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameAddrContext {
	var p = new(NameAddrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_nameAddr

	return p
}

func (s *NameAddrContext) GetParser() antlr.Parser { return s.parser }

func (s *NameAddrContext) AngleAddr() IAngleAddrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAngleAddrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAngleAddrContext)
}

func (s *NameAddrContext) DisplayName() IDisplayNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisplayNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisplayNameContext)
}

func (s *NameAddrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameAddrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameAddrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterNameAddr(s)
	}
}

func (s *NameAddrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitNameAddr(s)
	}
}

func (p *Address) NameAddr() (localctx INameAddrContext) {
	localctx = NewNameAddrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AddressRULE_nameAddr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(212)
			p.DisplayName()
		}

	}
	{
		p.SetState(215)
		p.AngleAddr()
	}

	return localctx
}

// IAngleAddrContext is an interface to support dynamic dispatch.
type IAngleAddrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAngleAddrContext differentiates from other interfaces.
	IsAngleAddrContext()
}

type AngleAddrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAngleAddrContext() *AngleAddrContext {
	var p = new(AngleAddrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_angleAddr
	return p
}

func (*AngleAddrContext) IsAngleAddrContext() {}

func NewAngleAddrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AngleAddrContext {
	var p = new(AngleAddrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_angleAddr

	return p
}

func (s *AngleAddrContext) GetParser() antlr.Parser { return s.parser }

func (s *AngleAddrContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressLess, 0)
}

func (s *AngleAddrContext) AddrSpec() IAddrSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddrSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddrSpecContext)
}

func (s *AngleAddrContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressGreater, 0)
}

func (s *AngleAddrContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *AngleAddrContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *AngleAddrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AngleAddrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AngleAddrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterAngleAddr(s)
	}
}

func (s *AngleAddrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitAngleAddr(s)
	}
}

func (p *Address) AngleAddr() (localctx IAngleAddrContext) {
	localctx = NewAngleAddrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AddressRULE_angleAddr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(217)
			p.Cfws()
		}

	}
	{
		p.SetState(220)
		p.Match(AddressLess)
	}
	{
		p.SetState(221)
		p.AddrSpec()
	}
	{
		p.SetState(222)
		p.Match(AddressGreater)
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(223)
			p.Cfws()
		}

	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) DisplayName() IDisplayNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisplayNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisplayNameContext)
}

func (s *GroupContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressColon, 0)
}

func (s *GroupContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressSemicolon, 0)
}

func (s *GroupContext) GroupList() IGroupListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupListContext)
}

func (s *GroupContext) Cfws() ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *Address) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AddressRULE_group)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.DisplayName()
	}
	{
		p.SetState(227)
		p.Match(AddressColon)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressExclamation)|(1<<AddressDQuote)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressLParens)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressMinus)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressLess)|(1<<AddressEqual)|(1<<AddressQuestion)|(1<<AddressAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0) {
		{
			p.SetState(228)
			p.GroupList()
		}

	}
	{
		p.SetState(231)
		p.Match(AddressSemicolon)
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(232)
			p.Cfws()
		}

	}

	return localctx
}

// IDisplayNameContext is an interface to support dynamic dispatch.
type IDisplayNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisplayNameContext differentiates from other interfaces.
	IsDisplayNameContext()
}

type DisplayNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisplayNameContext() *DisplayNameContext {
	var p = new(DisplayNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_displayName
	return p
}

func (*DisplayNameContext) IsDisplayNameContext() {}

func NewDisplayNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisplayNameContext {
	var p = new(DisplayNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_displayName

	return p
}

func (s *DisplayNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DisplayNameContext) Phrase() IPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPhraseContext)
}

func (s *DisplayNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisplayNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisplayNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterDisplayName(s)
	}
}

func (s *DisplayNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitDisplayName(s)
	}
}

func (p *Address) DisplayName() (localctx IDisplayNameContext) {
	localctx = NewDisplayNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AddressRULE_displayName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Phrase()
	}

	return localctx
}

// IMailboxListContext is an interface to support dynamic dispatch.
type IMailboxListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMailboxListContext differentiates from other interfaces.
	IsMailboxListContext()
}

type MailboxListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMailboxListContext() *MailboxListContext {
	var p = new(MailboxListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_mailboxList
	return p
}

func (*MailboxListContext) IsMailboxListContext() {}

func NewMailboxListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MailboxListContext {
	var p = new(MailboxListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_mailboxList

	return p
}

func (s *MailboxListContext) GetParser() antlr.Parser { return s.parser }

func (s *MailboxListContext) AllMailbox() []IMailboxContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMailboxContext)(nil)).Elem())
	var tst = make([]IMailboxContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMailboxContext)
		}
	}

	return tst
}

func (s *MailboxListContext) Mailbox(i int) IMailboxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMailboxContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMailboxContext)
}

func (s *MailboxListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AddressComma)
}

func (s *MailboxListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AddressComma, i)
}

func (s *MailboxListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MailboxListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MailboxListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterMailboxList(s)
	}
}

func (s *MailboxListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitMailboxList(s)
	}
}

func (p *Address) MailboxList() (localctx IMailboxListContext) {
	localctx = NewMailboxListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AddressRULE_mailboxList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Mailbox()
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressComma {
		{
			p.SetState(238)
			p.Match(AddressComma)
		}
		{
			p.SetState(239)
			p.Mailbox()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAddressListContext is an interface to support dynamic dispatch.
type IAddressListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddressListContext differentiates from other interfaces.
	IsAddressListContext()
}

type AddressListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddressListContext() *AddressListContext {
	var p = new(AddressListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_addressList
	return p
}

func (*AddressListContext) IsAddressListContext() {}

func NewAddressListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddressListContext {
	var p = new(AddressListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_addressList

	return p
}

func (s *AddressListContext) GetParser() antlr.Parser { return s.parser }

func (s *AddressListContext) AllAddress() []IAddressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddressContext)(nil)).Elem())
	var tst = make([]IAddressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddressContext)
		}
	}

	return tst
}

func (s *AddressListContext) Address(i int) IAddressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddressContext)
}

func (s *AddressListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AddressComma)
}

func (s *AddressListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AddressComma, i)
}

func (s *AddressListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddressListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddressListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterAddressList(s)
	}
}

func (s *AddressListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitAddressList(s)
	}
}

func (p *Address) AddressList() (localctx IAddressListContext) {
	localctx = NewAddressListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AddressRULE_addressList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Address()
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressComma {
		{
			p.SetState(246)
			p.Match(AddressComma)
		}
		{
			p.SetState(247)
			p.Address()
		}

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroupListContext is an interface to support dynamic dispatch.
type IGroupListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupListContext differentiates from other interfaces.
	IsGroupListContext()
}

type GroupListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupListContext() *GroupListContext {
	var p = new(GroupListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_groupList
	return p
}

func (*GroupListContext) IsGroupListContext() {}

func NewGroupListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupListContext {
	var p = new(GroupListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_groupList

	return p
}

func (s *GroupListContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupListContext) MailboxList() IMailboxListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMailboxListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMailboxListContext)
}

func (s *GroupListContext) Cfws() ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *GroupListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterGroupList(s)
	}
}

func (s *GroupListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitGroupList(s)
	}
}

func (p *Address) GroupList() (localctx IGroupListContext) {
	localctx = NewGroupListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AddressRULE_groupList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.MailboxList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Cfws()
		}

	}

	return localctx
}

// IAddrSpecContext is an interface to support dynamic dispatch.
type IAddrSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddrSpecContext differentiates from other interfaces.
	IsAddrSpecContext()
}

type AddrSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddrSpecContext() *AddrSpecContext {
	var p = new(AddrSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_addrSpec
	return p
}

func (*AddrSpecContext) IsAddrSpecContext() {}

func NewAddrSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddrSpecContext {
	var p = new(AddrSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_addrSpec

	return p
}

func (s *AddrSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *AddrSpecContext) LocalPart() ILocalPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalPartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalPartContext)
}

func (s *AddrSpecContext) At() antlr.TerminalNode {
	return s.GetToken(AddressAt, 0)
}

func (s *AddrSpecContext) Domain() IDomainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDomainContext)
}

func (s *AddrSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddrSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddrSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterAddrSpec(s)
	}
}

func (s *AddrSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitAddrSpec(s)
	}
}

func (p *Address) AddrSpec() (localctx IAddrSpecContext) {
	localctx = NewAddrSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AddressRULE_addrSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.LocalPart()
	}
	{
		p.SetState(258)
		p.Match(AddressAt)
	}
	{
		p.SetState(259)
		p.Domain()
	}

	return localctx
}

// ILocalPartContext is an interface to support dynamic dispatch.
type ILocalPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalPartContext differentiates from other interfaces.
	IsLocalPartContext()
}

type LocalPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalPartContext() *LocalPartContext {
	var p = new(LocalPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_localPart
	return p
}

func (*LocalPartContext) IsLocalPartContext() {}

func NewLocalPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalPartContext {
	var p = new(LocalPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_localPart

	return p
}

func (s *LocalPartContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalPartContext) DotAtom() IDotAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDotAtomContext)
}

func (s *LocalPartContext) QuotedString() IQuotedStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedStringContext)
}

func (s *LocalPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterLocalPart(s)
	}
}

func (s *LocalPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitLocalPart(s)
	}
}

func (p *Address) LocalPart() (localctx ILocalPartContext) {
	localctx = NewLocalPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AddressRULE_localPart)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.DotAtom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.QuotedString()
		}

	}

	return localctx
}

// IDomainContext is an interface to support dynamic dispatch.
type IDomainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDomainContext differentiates from other interfaces.
	IsDomainContext()
}

type DomainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainContext() *DomainContext {
	var p = new(DomainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_domain
	return p
}

func (*DomainContext) IsDomainContext() {}

func NewDomainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainContext {
	var p = new(DomainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_domain

	return p
}

func (s *DomainContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainContext) DotAtom() IDotAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDotAtomContext)
}

func (s *DomainContext) DomainLiteral() IDomainLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDomainLiteralContext)
}

func (s *DomainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterDomain(s)
	}
}

func (s *DomainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitDomain(s)
	}
}

func (p *Address) Domain() (localctx IDomainContext) {
	localctx = NewDomainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AddressRULE_domain)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.DotAtom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.DomainLiteral()
		}

	}

	return localctx
}

// IDomainLiteralContext is an interface to support dynamic dispatch.
type IDomainLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDomainLiteralContext differentiates from other interfaces.
	IsDomainLiteralContext()
}

type DomainLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainLiteralContext() *DomainLiteralContext {
	var p = new(DomainLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_domainLiteral
	return p
}

func (*DomainLiteralContext) IsDomainLiteralContext() {}

func NewDomainLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainLiteralContext {
	var p = new(DomainLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_domainLiteral

	return p
}

func (s *DomainLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainLiteralContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressLBracket, 0)
}

func (s *DomainLiteralContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressRBracket, 0)
}

func (s *DomainLiteralContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *DomainLiteralContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *DomainLiteralContext) AllDtext() []IDtextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDtextContext)(nil)).Elem())
	var tst = make([]IDtextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDtextContext)
		}
	}

	return tst
}

func (s *DomainLiteralContext) Dtext(i int) IDtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDtextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDtextContext)
}

func (s *DomainLiteralContext) AllFws() []IFwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFwsContext)(nil)).Elem())
	var tst = make([]IFwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFwsContext)
		}
	}

	return tst
}

func (s *DomainLiteralContext) Fws(i int) IFwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFwsContext)
}

func (s *DomainLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterDomainLiteral(s)
	}
}

func (s *DomainLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitDomainLiteral(s)
	}
}

func (p *Address) DomainLiteral() (localctx IDomainLiteralContext) {
	localctx = NewDomainLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AddressRULE_domainLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(269)
			p.Cfws()
		}

	}
	{
		p.SetState(272)
		p.Match(AddressLBracket)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressExclamation)|(1<<AddressDQuote)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressLParens)|(1<<AddressRParens)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressComma)|(1<<AddressMinus)|(1<<AddressPeriod)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressColon)|(1<<AddressSemicolon)|(1<<AddressLess)|(1<<AddressEqual)|(1<<AddressGreater)|(1<<AddressQuestion)|(1<<AddressAt)|(1<<AddressAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0) {
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP))) != 0 {
			{
				p.SetState(273)
				p.Fws()
			}

		}
		{
			p.SetState(276)
			p.Dtext()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(282)
		p.Match(AddressRBracket)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressCRLF)|(1<<AddressHTAB)|(1<<AddressSP)|(1<<AddressLParens))) != 0 {
		{
			p.SetState(283)
			p.Cfws()
		}

	}

	return localctx
}

// IDtextContext is an interface to support dynamic dispatch.
type IDtextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDtextContext differentiates from other interfaces.
	IsDtextContext()
}

type DtextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDtextContext() *DtextContext {
	var p = new(DtextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_dtext
	return p
}

func (*DtextContext) IsDtextContext() {}

func NewDtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DtextContext {
	var p = new(DtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_dtext

	return p
}

func (s *DtextContext) GetParser() antlr.Parser { return s.parser }

func (s *DtextContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressExclamation, 0)
}

func (s *DtextContext) DQuote() antlr.TerminalNode {
	return s.GetToken(AddressDQuote, 0)
}

func (s *DtextContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressHash, 0)
}

func (s *DtextContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressDollar, 0)
}

func (s *DtextContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressPercent, 0)
}

func (s *DtextContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressAmpersand, 0)
}

func (s *DtextContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressSQuote, 0)
}

func (s *DtextContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressLParens, 0)
}

func (s *DtextContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressRParens, 0)
}

func (s *DtextContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressAsterisk, 0)
}

func (s *DtextContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressPlus, 0)
}

func (s *DtextContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressComma, 0)
}

func (s *DtextContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressMinus, 0)
}

func (s *DtextContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressPeriod, 0)
}

func (s *DtextContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressSlash, 0)
}

func (s *DtextContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressDigit, 0)
}

func (s *DtextContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressColon, 0)
}

func (s *DtextContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressSemicolon, 0)
}

func (s *DtextContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressLess, 0)
}

func (s *DtextContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressEqual, 0)
}

func (s *DtextContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressGreater, 0)
}

func (s *DtextContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressQuestion, 0)
}

func (s *DtextContext) At() antlr.TerminalNode {
	return s.GetToken(AddressAt, 0)
}

func (s *DtextContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressAlphaUpper, 0)
}

func (s *DtextContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressCaret, 0)
}

func (s *DtextContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressUnderscore, 0)
}

func (s *DtextContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressBacktick, 0)
}

func (s *DtextContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressAlphaLower, 0)
}

func (s *DtextContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressLCurly, 0)
}

func (s *DtextContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressPipe, 0)
}

func (s *DtextContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressRCurly, 0)
}

func (s *DtextContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressTilde, 0)
}

func (s *DtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterDtext(s)
	}
}

func (s *DtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitDtext(s)
	}
}

func (p *Address) Dtext() (localctx IDtextContext) {
	localctx = NewDtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, AddressRULE_dtext)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressExclamation)|(1<<AddressDQuote)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressLParens)|(1<<AddressRParens)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressComma)|(1<<AddressMinus)|(1<<AddressPeriod)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressColon)|(1<<AddressSemicolon)|(1<<AddressLess)|(1<<AddressEqual)|(1<<AddressGreater)|(1<<AddressQuestion)|(1<<AddressAt)|(1<<AddressAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWspContext is an interface to support dynamic dispatch.
type IWspContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWspContext differentiates from other interfaces.
	IsWspContext()
}

type WspContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWspContext() *WspContext {
	var p = new(WspContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_wsp
	return p
}

func (*WspContext) IsWspContext() {}

func NewWspContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WspContext {
	var p = new(WspContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_wsp

	return p
}

func (s *WspContext) GetParser() antlr.Parser { return s.parser }

func (s *WspContext) SP() antlr.TerminalNode {
	return s.GetToken(AddressSP, 0)
}

func (s *WspContext) HTAB() antlr.TerminalNode {
	return s.GetToken(AddressHTAB, 0)
}

func (s *WspContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WspContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WspContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterWsp(s)
	}
}

func (s *WspContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitWsp(s)
	}
}

func (p *Address) Wsp() (localctx IWspContext) {
	localctx = NewWspContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AddressRULE_wsp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AddressHTAB || _la == AddressSP) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVcharContext is an interface to support dynamic dispatch.
type IVcharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVcharContext differentiates from other interfaces.
	IsVcharContext()
}

type VcharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVcharContext() *VcharContext {
	var p = new(VcharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressRULE_vchar
	return p
}

func (*VcharContext) IsVcharContext() {}

func NewVcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VcharContext {
	var p = new(VcharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressRULE_vchar

	return p
}

func (s *VcharContext) GetParser() antlr.Parser { return s.parser }

func (s *VcharContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressExclamation, 0)
}

func (s *VcharContext) DQuote() antlr.TerminalNode {
	return s.GetToken(AddressDQuote, 0)
}

func (s *VcharContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressHash, 0)
}

func (s *VcharContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressDollar, 0)
}

func (s *VcharContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressPercent, 0)
}

func (s *VcharContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressAmpersand, 0)
}

func (s *VcharContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressSQuote, 0)
}

func (s *VcharContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressLParens, 0)
}

func (s *VcharContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressRParens, 0)
}

func (s *VcharContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressAsterisk, 0)
}

func (s *VcharContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressPlus, 0)
}

func (s *VcharContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressComma, 0)
}

func (s *VcharContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressMinus, 0)
}

func (s *VcharContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressPeriod, 0)
}

func (s *VcharContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressSlash, 0)
}

func (s *VcharContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressDigit, 0)
}

func (s *VcharContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressColon, 0)
}

func (s *VcharContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressSemicolon, 0)
}

func (s *VcharContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressLess, 0)
}

func (s *VcharContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressEqual, 0)
}

func (s *VcharContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressGreater, 0)
}

func (s *VcharContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressQuestion, 0)
}

func (s *VcharContext) At() antlr.TerminalNode {
	return s.GetToken(AddressAt, 0)
}

func (s *VcharContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressAlphaUpper, 0)
}

func (s *VcharContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressLBracket, 0)
}

func (s *VcharContext) Backslash() antlr.TerminalNode {
	return s.GetToken(AddressBackslash, 0)
}

func (s *VcharContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressRBracket, 0)
}

func (s *VcharContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressCaret, 0)
}

func (s *VcharContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressUnderscore, 0)
}

func (s *VcharContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressBacktick, 0)
}

func (s *VcharContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressAlphaLower, 0)
}

func (s *VcharContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressLCurly, 0)
}

func (s *VcharContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressPipe, 0)
}

func (s *VcharContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressRCurly, 0)
}

func (s *VcharContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressTilde, 0)
}

func (s *VcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.EnterVchar(s)
	}
}

func (s *VcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressListener); ok {
		listenerT.ExitVchar(s)
	}
}

func (p *Address) Vchar() (localctx IVcharContext) {
	localctx = NewVcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AddressRULE_vchar)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressExclamation)|(1<<AddressDQuote)|(1<<AddressHash)|(1<<AddressDollar)|(1<<AddressPercent)|(1<<AddressAmpersand)|(1<<AddressSQuote)|(1<<AddressLParens)|(1<<AddressRParens)|(1<<AddressAsterisk)|(1<<AddressPlus)|(1<<AddressComma)|(1<<AddressMinus)|(1<<AddressPeriod)|(1<<AddressSlash)|(1<<AddressDigit)|(1<<AddressColon)|(1<<AddressSemicolon)|(1<<AddressLess)|(1<<AddressEqual)|(1<<AddressGreater)|(1<<AddressQuestion)|(1<<AddressAt)|(1<<AddressAlphaUpper)|(1<<AddressLBracket)|(1<<AddressBackslash)|(1<<AddressRBracket))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressCaret-32))|(1<<(AddressUnderscore-32))|(1<<(AddressBacktick-32))|(1<<(AddressAlphaLower-32))|(1<<(AddressLCurly-32))|(1<<(AddressPipe-32))|(1<<(AddressRCurly-32))|(1<<(AddressTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
