// Code generated from pkg/message/address/parser/AddressParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AddressParser

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 302,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 5, 2, 75, 10, 2, 3, 3, 3,
	3, 3, 3, 3, 4, 7, 4, 81, 10, 4, 12, 4, 14, 4, 84, 11, 4, 3, 4, 5, 4, 87,
	10, 4, 3, 4, 6, 4, 90, 10, 4, 13, 4, 14, 4, 91, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 5, 6, 99, 10, 6, 3, 7, 3, 7, 5, 7, 103, 10, 7, 3, 7, 7, 7, 106, 10,
	7, 12, 7, 14, 7, 109, 11, 7, 3, 7, 5, 7, 112, 10, 7, 3, 7, 3, 7, 3, 8,
	5, 8, 117, 10, 8, 3, 8, 6, 8, 120, 10, 8, 13, 8, 14, 8, 121, 3, 8, 5, 8,
	125, 10, 8, 3, 8, 5, 8, 128, 10, 8, 3, 9, 3, 9, 3, 10, 6, 10, 133, 10,
	10, 13, 10, 14, 10, 134, 3, 11, 5, 11, 138, 10, 11, 3, 11, 3, 11, 5, 11,
	142, 10, 11, 3, 12, 3, 12, 3, 12, 7, 12, 147, 10, 12, 12, 12, 14, 12, 150,
	11, 12, 3, 13, 5, 13, 153, 10, 13, 3, 13, 3, 13, 5, 13, 157, 10, 13, 3,
	14, 3, 14, 3, 15, 3, 15, 5, 15, 163, 10, 15, 3, 16, 5, 16, 166, 10, 16,
	3, 16, 7, 16, 169, 10, 16, 12, 16, 14, 16, 172, 11, 16, 3, 17, 5, 17, 175,
	10, 17, 3, 17, 3, 17, 3, 17, 5, 17, 180, 10, 17, 3, 17, 3, 17, 5, 17, 184,
	10, 17, 3, 18, 3, 18, 5, 18, 188, 10, 18, 3, 19, 5, 19, 191, 10, 19, 3,
	19, 7, 19, 194, 10, 19, 12, 19, 14, 19, 197, 11, 19, 3, 19, 7, 19, 200,
	10, 19, 12, 19, 14, 19, 203, 11, 19, 3, 20, 3, 20, 5, 20, 207, 10, 20,
	3, 21, 3, 21, 5, 21, 211, 10, 21, 3, 22, 5, 22, 214, 10, 22, 3, 22, 3,
	22, 3, 23, 5, 23, 219, 10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 225,
	10, 23, 3, 24, 3, 24, 3, 24, 5, 24, 230, 10, 24, 3, 24, 3, 24, 5, 24, 234,
	10, 24, 3, 25, 6, 25, 237, 10, 25, 13, 25, 14, 25, 238, 3, 26, 3, 26, 3,
	26, 7, 26, 244, 10, 26, 12, 26, 14, 26, 247, 11, 26, 3, 27, 3, 27, 3, 27,
	7, 27, 252, 10, 27, 12, 27, 14, 27, 255, 11, 27, 3, 28, 3, 28, 5, 28, 259,
	10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 5, 30, 267, 10, 30, 3,
	31, 3, 31, 5, 31, 271, 10, 31, 3, 32, 5, 32, 274, 10, 32, 3, 32, 3, 32,
	5, 32, 278, 10, 32, 3, 32, 7, 32, 281, 10, 32, 12, 32, 14, 32, 284, 11,
	32, 3, 32, 5, 32, 287, 10, 32, 3, 32, 3, 32, 5, 32, 291, 10, 32, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 2, 2, 37,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 2, 8, 5,
	2, 7, 13, 16, 31, 33, 41, 11, 2, 7, 7, 9, 13, 16, 17, 19, 19, 21, 22, 26,
	26, 28, 28, 30, 30, 34, 41, 5, 2, 7, 7, 9, 31, 33, 41, 4, 2, 7, 30, 34,
	41, 4, 2, 3, 3, 6, 6, 3, 2, 7, 41, 2, 313, 2, 74, 3, 2, 2, 2, 4, 76, 3,
	2, 2, 2, 6, 86, 3, 2, 2, 2, 8, 93, 3, 2, 2, 2, 10, 98, 3, 2, 2, 2, 12,
	100, 3, 2, 2, 2, 14, 127, 3, 2, 2, 2, 16, 129, 3, 2, 2, 2, 18, 132, 3,
	2, 2, 2, 20, 137, 3, 2, 2, 2, 22, 143, 3, 2, 2, 2, 24, 152, 3, 2, 2, 2,
	26, 158, 3, 2, 2, 2, 28, 162, 3, 2, 2, 2, 30, 170, 3, 2, 2, 2, 32, 174,
	3, 2, 2, 2, 34, 187, 3, 2, 2, 2, 36, 195, 3, 2, 2, 2, 38, 206, 3, 2, 2,
	2, 40, 210, 3, 2, 2, 2, 42, 213, 3, 2, 2, 2, 44, 218, 3, 2, 2, 2, 46, 226,
	3, 2, 2, 2, 48, 236, 3, 2, 2, 2, 50, 240, 3, 2, 2, 2, 52, 248, 3, 2, 2,
	2, 54, 258, 3, 2, 2, 2, 56, 260, 3, 2, 2, 2, 58, 266, 3, 2, 2, 2, 60, 270,
	3, 2, 2, 2, 62, 273, 3, 2, 2, 2, 64, 292, 3, 2, 2, 2, 66, 294, 3, 2, 2,
	2, 68, 297, 3, 2, 2, 2, 70, 299, 3, 2, 2, 2, 72, 75, 5, 70, 36, 2, 73,
	75, 5, 68, 35, 2, 74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 3, 3, 2,
	2, 2, 76, 77, 7, 32, 2, 2, 77, 78, 5, 2, 2, 2, 78, 5, 3, 2, 2, 2, 79, 81,
	5, 68, 35, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2,
	2, 82, 83, 3, 2, 2, 2, 83, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 87,
	5, 66, 34, 2, 86, 82, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2,
	2, 88, 90, 5, 68, 35, 2, 89, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 89,
	3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 7, 3, 2, 2, 2, 93, 94, 9, 2, 2, 2,
	94, 9, 3, 2, 2, 2, 95, 99, 5, 8, 5, 2, 96, 99, 5, 4, 3, 2, 97, 99, 5, 12,
	7, 2, 98, 95, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 97, 3, 2, 2, 2, 99, 11,
	3, 2, 2, 2, 100, 107, 7, 14, 2, 2, 101, 103, 5, 6, 4, 2, 102, 101, 3, 2,
	2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 106, 5, 10, 6, 2,
	105, 102, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107,
	108, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 112,
	5, 6, 4, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 3, 2,
	2, 2, 113, 114, 7, 15, 2, 2, 114, 13, 3, 2, 2, 2, 115, 117, 5, 6, 4, 2,
	116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118,
	120, 5, 12, 7, 2, 119, 116, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 119,
	3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 124, 3, 2, 2, 2, 123, 125, 5, 6,
	4, 2, 124, 123, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2,
	126, 128, 5, 6, 4, 2, 127, 119, 3, 2, 2, 2, 127, 126, 3, 2, 2, 2, 128,
	15, 3, 2, 2, 2, 129, 130, 9, 3, 2, 2, 130, 17, 3, 2, 2, 2, 131, 133, 5,
	16, 9, 2, 132, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 132, 3, 2, 2,
	2, 134, 135, 3, 2, 2, 2, 135, 19, 3, 2, 2, 2, 136, 138, 5, 14, 8, 2, 137,
	136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 141,
	5, 18, 10, 2, 140, 142, 5, 14, 8, 2, 141, 140, 3, 2, 2, 2, 141, 142, 3,
	2, 2, 2, 142, 21, 3, 2, 2, 2, 143, 148, 5, 18, 10, 2, 144, 145, 7, 20,
	2, 2, 145, 147, 5, 18, 10, 2, 146, 144, 3, 2, 2, 2, 147, 150, 3, 2, 2,
	2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 23, 3, 2, 2, 2, 150,
	148, 3, 2, 2, 2, 151, 153, 5, 14, 8, 2, 152, 151, 3, 2, 2, 2, 152, 153,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 156, 5, 22, 12, 2, 155, 157, 5,
	14, 8, 2, 156, 155, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 25, 3, 2, 2,
	2, 158, 159, 9, 4, 2, 2, 159, 27, 3, 2, 2, 2, 160, 163, 5, 26, 14, 2, 161,
	163, 5, 4, 3, 2, 162, 160, 3, 2, 2, 2, 162, 161, 3, 2, 2, 2, 163, 29, 3,
	2, 2, 2, 164, 166, 5, 6, 4, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2,
	2, 166, 167, 3, 2, 2, 2, 167, 169, 5, 28, 15, 2, 168, 165, 3, 2, 2, 2,
	169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171,
	31, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 175, 5, 14, 8, 2, 174, 173,
	3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 177, 7, 8,
	2, 2, 177, 179, 5, 30, 16, 2, 178, 180, 5, 6, 4, 2, 179, 178, 3, 2, 2,
	2, 179, 180, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 7, 8, 2, 2, 182,
	184, 5, 14, 8, 2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 33,
	3, 2, 2, 2, 185, 188, 5, 20, 11, 2, 186, 188, 5, 32, 17, 2, 187, 185, 3,
	2, 2, 2, 187, 186, 3, 2, 2, 2, 188, 35, 3, 2, 2, 2, 189, 191, 5, 6, 4,
	2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192,
	194, 5, 70, 36, 2, 193, 190, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193,
	3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 201, 3, 2, 2, 2, 197, 195, 3, 2,
	2, 2, 198, 200, 5, 68, 35, 2, 199, 198, 3, 2, 2, 2, 200, 203, 3, 2, 2,
	2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 37, 3, 2, 2, 2, 203,
	201, 3, 2, 2, 2, 204, 207, 5, 40, 21, 2, 205, 207, 5, 46, 24, 2, 206, 204,
	3, 2, 2, 2, 206, 205, 3, 2, 2, 2, 207, 39, 3, 2, 2, 2, 208, 211, 5, 42,
	22, 2, 209, 211, 5, 56, 29, 2, 210, 208, 3, 2, 2, 2, 210, 209, 3, 2, 2,
	2, 211, 41, 3, 2, 2, 2, 212, 214, 5, 48, 25, 2, 213, 212, 3, 2, 2, 2, 213,
	214, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 5, 44, 23, 2, 216, 43,
	3, 2, 2, 2, 217, 219, 5, 14, 8, 2, 218, 217, 3, 2, 2, 2, 218, 219, 3, 2,
	2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 7, 25, 2, 2, 221, 222, 5, 56, 29,
	2, 222, 224, 7, 27, 2, 2, 223, 225, 5, 14, 8, 2, 224, 223, 3, 2, 2, 2,
	224, 225, 3, 2, 2, 2, 225, 45, 3, 2, 2, 2, 226, 227, 5, 48, 25, 2, 227,
	229, 7, 23, 2, 2, 228, 230, 5, 54, 28, 2, 229, 228, 3, 2, 2, 2, 229, 230,
	3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 233, 7, 24, 2, 2, 232, 234, 5, 14,
	8, 2, 233, 232, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 47, 3, 2, 2, 2,
	235, 237, 5, 34, 18, 2, 236, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238,
	236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 49, 3, 2, 2, 2, 240, 245, 5,
	40, 21, 2, 241, 242, 7, 18, 2, 2, 242, 244, 5, 40, 21, 2, 243, 241, 3,
	2, 2, 2, 244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2,
	2, 246, 51, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 253, 5, 38, 20, 2, 249,
	250, 7, 18, 2, 2, 250, 252, 5, 38, 20, 2, 251, 249, 3, 2, 2, 2, 252, 255,
	3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 53, 3, 2,
	2, 2, 255, 253, 3, 2, 2, 2, 256, 259, 5, 50, 26, 2, 257, 259, 5, 14, 8,
	2, 258, 256, 3, 2, 2, 2, 258, 257, 3, 2, 2, 2, 259, 55, 3, 2, 2, 2, 260,
	261, 5, 58, 30, 2, 261, 262, 7, 29, 2, 2, 262, 263, 5, 60, 31, 2, 263,
	57, 3, 2, 2, 2, 264, 267, 5, 24, 13, 2, 265, 267, 5, 32, 17, 2, 266, 264,
	3, 2, 2, 2, 266, 265, 3, 2, 2, 2, 267, 59, 3, 2, 2, 2, 268, 271, 5, 24,
	13, 2, 269, 271, 5, 62, 32, 2, 270, 268, 3, 2, 2, 2, 270, 269, 3, 2, 2,
	2, 271, 61, 3, 2, 2, 2, 272, 274, 5, 14, 8, 2, 273, 272, 3, 2, 2, 2, 273,
	274, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 282, 7, 31, 2, 2, 276, 278,
	5, 6, 4, 2, 277, 276, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 3, 2,
	2, 2, 279, 281, 5, 64, 33, 2, 280, 277, 3, 2, 2, 2, 281, 284, 3, 2, 2,
	2, 282, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 286, 3, 2, 2, 2, 284,
	282, 3, 2, 2, 2, 285, 287, 5, 6, 4, 2, 286, 285, 3, 2, 2, 2, 286, 287,
	3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 290, 7, 33, 2, 2, 289, 291, 5, 14,
	8, 2, 290, 289, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 63, 3, 2, 2, 2,
	292, 293, 9, 5, 2, 2, 293, 65, 3, 2, 2, 2, 294, 295, 7, 5, 2, 2, 295, 296,
	7, 4, 2, 2, 296, 67, 3, 2, 2, 2, 297, 298, 9, 6, 2, 2, 298, 69, 3, 2, 2,
	2, 299, 300, 9, 7, 2, 2, 300, 71, 3, 2, 2, 2, 48, 74, 82, 86, 91, 98, 102,
	107, 111, 116, 121, 124, 127, 134, 137, 141, 148, 152, 156, 162, 165, 170,
	174, 179, 183, 187, 190, 195, 201, 206, 210, 213, 218, 224, 229, 233, 238,
	245, 253, 258, 266, 270, 273, 277, 282, 286, 290,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\t'", "'\n'", "'\r'", "' '", "'!'", "'\"'", "'#'", "'$'", "'%'",
	"'&'", "'''", "'('", "')'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "",
	"':'", "';'", "'<'", "'='", "'>'", "'?'", "'@'", "", "'['", "'\\'", "']'",
	"'^'", "'_'", "'`'", "", "'{'", "'|'", "'}'", "'~'", "'\u007F'",
}
var symbolicNames = []string{
	"", "TAB", "LF", "CR", "SP", "Exclamation", "DQuote", "Hash", "Dollar",
	"Percent", "Ampersand", "SQuote", "LParens", "RParens", "Asterisk", "Plus",
	"Comma", "Minus", "Period", "Slash", "Digit", "Colon", "Semicolon", "Less",
	"Equal", "Greater", "Question", "At", "AlphaUpper", "LBracket", "Backslash",
	"RBracket", "Caret", "Underscore", "Backtick", "AlphaLower", "LCurly",
	"Pipe", "RCurly", "Tilde", "DEL",
}

var ruleNames = []string{
	"quotedChar", "quotedPair", "fws", "ctext", "ccontent", "comment", "cfws",
	"atext", "atextString", "atom", "dotAtomText", "dotAtom", "qtext", "quotedContent",
	"quotedValue", "quotedString", "word", "unstructured", "address", "mailbox",
	"nameAddr", "angleAddr", "group", "displayName", "mailboxList", "addressList",
	"groupList", "addrSpec", "localPart", "domain", "domainLiteral", "dtext",
	"crlf", "wsp", "vchar",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AddressParser struct {
	*antlr.BaseParser
}

func NewAddressParser(input antlr.TokenStream) *AddressParser {
	this := new(AddressParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AddressParser.g4"

	return this
}

// AddressParser tokens.
const (
	AddressParserEOF         = antlr.TokenEOF
	AddressParserTAB         = 1
	AddressParserLF          = 2
	AddressParserCR          = 3
	AddressParserSP          = 4
	AddressParserExclamation = 5
	AddressParserDQuote      = 6
	AddressParserHash        = 7
	AddressParserDollar      = 8
	AddressParserPercent     = 9
	AddressParserAmpersand   = 10
	AddressParserSQuote      = 11
	AddressParserLParens     = 12
	AddressParserRParens     = 13
	AddressParserAsterisk    = 14
	AddressParserPlus        = 15
	AddressParserComma       = 16
	AddressParserMinus       = 17
	AddressParserPeriod      = 18
	AddressParserSlash       = 19
	AddressParserDigit       = 20
	AddressParserColon       = 21
	AddressParserSemicolon   = 22
	AddressParserLess        = 23
	AddressParserEqual       = 24
	AddressParserGreater     = 25
	AddressParserQuestion    = 26
	AddressParserAt          = 27
	AddressParserAlphaUpper  = 28
	AddressParserLBracket    = 29
	AddressParserBackslash   = 30
	AddressParserRBracket    = 31
	AddressParserCaret       = 32
	AddressParserUnderscore  = 33
	AddressParserBacktick    = 34
	AddressParserAlphaLower  = 35
	AddressParserLCurly      = 36
	AddressParserPipe        = 37
	AddressParserRCurly      = 38
	AddressParserTilde       = 39
	AddressParserDEL         = 40
)

// AddressParser rules.
const (
	AddressParserRULE_quotedChar    = 0
	AddressParserRULE_quotedPair    = 1
	AddressParserRULE_fws           = 2
	AddressParserRULE_ctext         = 3
	AddressParserRULE_ccontent      = 4
	AddressParserRULE_comment       = 5
	AddressParserRULE_cfws          = 6
	AddressParserRULE_atext         = 7
	AddressParserRULE_atextString   = 8
	AddressParserRULE_atom          = 9
	AddressParserRULE_dotAtomText   = 10
	AddressParserRULE_dotAtom       = 11
	AddressParserRULE_qtext         = 12
	AddressParserRULE_quotedContent = 13
	AddressParserRULE_quotedValue   = 14
	AddressParserRULE_quotedString  = 15
	AddressParserRULE_word          = 16
	AddressParserRULE_unstructured  = 17
	AddressParserRULE_address       = 18
	AddressParserRULE_mailbox       = 19
	AddressParserRULE_nameAddr      = 20
	AddressParserRULE_angleAddr     = 21
	AddressParserRULE_group         = 22
	AddressParserRULE_displayName   = 23
	AddressParserRULE_mailboxList   = 24
	AddressParserRULE_addressList   = 25
	AddressParserRULE_groupList     = 26
	AddressParserRULE_addrSpec      = 27
	AddressParserRULE_localPart     = 28
	AddressParserRULE_domain        = 29
	AddressParserRULE_domainLiteral = 30
	AddressParserRULE_dtext         = 31
	AddressParserRULE_crlf          = 32
	AddressParserRULE_wsp           = 33
	AddressParserRULE_vchar         = 34
)

// IQuotedCharContext is an interface to support dynamic dispatch.
type IQuotedCharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedCharContext differentiates from other interfaces.
	IsQuotedCharContext()
}

type QuotedCharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedCharContext() *QuotedCharContext {
	var p = new(QuotedCharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_quotedChar
	return p
}

func (*QuotedCharContext) IsQuotedCharContext() {}

func NewQuotedCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedCharContext {
	var p = new(QuotedCharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_quotedChar

	return p
}

func (s *QuotedCharContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedCharContext) Vchar() IVcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVcharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVcharContext)
}

func (s *QuotedCharContext) Wsp() IWspContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWspContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWspContext)
}

func (s *QuotedCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedCharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterQuotedChar(s)
	}
}

func (s *QuotedCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitQuotedChar(s)
	}
}

func (p *AddressParser) QuotedChar() (localctx IQuotedCharContext) {
	localctx = NewQuotedCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AddressParserRULE_quotedChar)

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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation, AddressParserDQuote, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserLParens, AddressParserRParens, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserBackslash, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Vchar()
		}

	case AddressParserTAB, AddressParserSP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Wsp()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

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
	p.RuleIndex = AddressParserRULE_quotedPair
	return p
}

func (*QuotedPairContext) IsQuotedPairContext() {}

func NewQuotedPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedPairContext {
	var p = new(QuotedPairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_quotedPair

	return p
}

func (s *QuotedPairContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedPairContext) Backslash() antlr.TerminalNode {
	return s.GetToken(AddressParserBackslash, 0)
}

func (s *QuotedPairContext) QuotedChar() IQuotedCharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedCharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedCharContext)
}

func (s *QuotedPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterQuotedPair(s)
	}
}

func (s *QuotedPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitQuotedPair(s)
	}
}

func (p *AddressParser) QuotedPair() (localctx IQuotedPairContext) {
	localctx = NewQuotedPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AddressParserRULE_quotedPair)

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
		p.SetState(74)
		p.Match(AddressParserBackslash)
	}
	{
		p.SetState(75)
		p.QuotedChar()
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
	p.RuleIndex = AddressParserRULE_fws
	return p
}

func (*FwsContext) IsFwsContext() {}

func NewFwsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FwsContext {
	var p = new(FwsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_fws

	return p
}

func (s *FwsContext) GetParser() antlr.Parser { return s.parser }

func (s *FwsContext) Crlf() ICrlfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICrlfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICrlfContext)
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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterFws(s)
	}
}

func (s *FwsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitFws(s)
	}
}

func (p *AddressParser) Fws() (localctx IFwsContext) {
	localctx = NewFwsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AddressParserRULE_fws)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AddressParserTAB || _la == AddressParserSP {
			{
				p.SetState(77)
				p.Wsp()
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(83)
			p.Crlf()
		}

	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(86)
				p.Wsp()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.RuleIndex = AddressParserRULE_ctext
	return p
}

func (*CtextContext) IsCtextContext() {}

func NewCtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CtextContext {
	var p = new(CtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_ctext

	return p
}

func (s *CtextContext) GetParser() antlr.Parser { return s.parser }

func (s *CtextContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressParserExclamation, 0)
}

func (s *CtextContext) DQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserDQuote, 0)
}

func (s *CtextContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressParserHash, 0)
}

func (s *CtextContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressParserDollar, 0)
}

func (s *CtextContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressParserPercent, 0)
}

func (s *CtextContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressParserAmpersand, 0)
}

func (s *CtextContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserSQuote, 0)
}

func (s *CtextContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressParserAsterisk, 0)
}

func (s *CtextContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressParserPlus, 0)
}

func (s *CtextContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressParserComma, 0)
}

func (s *CtextContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressParserMinus, 0)
}

func (s *CtextContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, 0)
}

func (s *CtextContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressParserSlash, 0)
}

func (s *CtextContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressParserDigit, 0)
}

func (s *CtextContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressParserColon, 0)
}

func (s *CtextContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressParserSemicolon, 0)
}

func (s *CtextContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressParserLess, 0)
}

func (s *CtextContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressParserEqual, 0)
}

func (s *CtextContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressParserGreater, 0)
}

func (s *CtextContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressParserQuestion, 0)
}

func (s *CtextContext) At() antlr.TerminalNode {
	return s.GetToken(AddressParserAt, 0)
}

func (s *CtextContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaUpper, 0)
}

func (s *CtextContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserLBracket, 0)
}

func (s *CtextContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserRBracket, 0)
}

func (s *CtextContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressParserCaret, 0)
}

func (s *CtextContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressParserUnderscore, 0)
}

func (s *CtextContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressParserBacktick, 0)
}

func (s *CtextContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaLower, 0)
}

func (s *CtextContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserLCurly, 0)
}

func (s *CtextContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressParserPipe, 0)
}

func (s *CtextContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserRCurly, 0)
}

func (s *CtextContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressParserTilde, 0)
}

func (s *CtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterCtext(s)
	}
}

func (s *CtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitCtext(s)
	}
}

func (p *AddressParser) Ctext() (localctx ICtextContext) {
	localctx = NewCtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AddressParserRULE_ctext)
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
		p.SetState(91)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater)|(1<<AddressParserQuestion)|(1<<AddressParserAt)|(1<<AddressParserAlphaUpper)|(1<<AddressParserLBracket)|(1<<AddressParserRBracket))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0)) {
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
	p.RuleIndex = AddressParserRULE_ccontent
	return p
}

func (*CcontentContext) IsCcontentContext() {}

func NewCcontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CcontentContext {
	var p = new(CcontentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_ccontent

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterCcontent(s)
	}
}

func (s *CcontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitCcontent(s)
	}
}

func (p *AddressParser) Ccontent() (localctx ICcontentContext) {
	localctx = NewCcontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AddressParserRULE_ccontent)

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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation, AddressParserDQuote, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Ctext()
		}

	case AddressParserBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.QuotedPair()
		}

	case AddressParserLParens:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
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
	p.RuleIndex = AddressParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressParserLParens, 0)
}

func (s *CommentContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressParserRParens, 0)
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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *AddressParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AddressParserRULE_comment)
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
		p.SetState(98)
		p.Match(AddressParserLParens)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(99)
					p.Fws()
				}

			}
			{
				p.SetState(102)
				p.Ccontent()
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(108)
			p.Fws()
		}

	}
	{
		p.SetState(111)
		p.Match(AddressParserRParens)
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
	p.RuleIndex = AddressParserRULE_cfws
	return p
}

func (*CfwsContext) IsCfwsContext() {}

func NewCfwsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CfwsContext {
	var p = new(CfwsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_cfws

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterCfws(s)
	}
}

func (s *CfwsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitCfws(s)
	}
}

func (p *AddressParser) Cfws() (localctx ICfwsContext) {
	localctx = NewCfwsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AddressParserRULE_cfws)
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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(114)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
					{
						p.SetState(113)
						p.Fws()
					}

				}
				{
					p.SetState(116)
					p.Comment()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(121)
				p.Fws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
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
	p.RuleIndex = AddressParserRULE_atext
	return p
}

func (*AtextContext) IsAtextContext() {}

func NewAtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtextContext {
	var p = new(AtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_atext

	return p
}

func (s *AtextContext) GetParser() antlr.Parser { return s.parser }

func (s *AtextContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaUpper, 0)
}

func (s *AtextContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaLower, 0)
}

func (s *AtextContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressParserDigit, 0)
}

func (s *AtextContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressParserExclamation, 0)
}

func (s *AtextContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressParserHash, 0)
}

func (s *AtextContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressParserDollar, 0)
}

func (s *AtextContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressParserPercent, 0)
}

func (s *AtextContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressParserAmpersand, 0)
}

func (s *AtextContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserSQuote, 0)
}

func (s *AtextContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressParserAsterisk, 0)
}

func (s *AtextContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressParserPlus, 0)
}

func (s *AtextContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressParserMinus, 0)
}

func (s *AtextContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressParserSlash, 0)
}

func (s *AtextContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressParserEqual, 0)
}

func (s *AtextContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressParserQuestion, 0)
}

func (s *AtextContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressParserCaret, 0)
}

func (s *AtextContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressParserUnderscore, 0)
}

func (s *AtextContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressParserBacktick, 0)
}

func (s *AtextContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserLCurly, 0)
}

func (s *AtextContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressParserPipe, 0)
}

func (s *AtextContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserRCurly, 0)
}

func (s *AtextContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressParserTilde, 0)
}

func (s *AtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterAtext(s)
	}
}

func (s *AtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitAtext(s)
	}
}

func (p *AddressParser) Atext() (localctx IAtextContext) {
	localctx = NewAtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AddressParserRULE_atext)
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
		p.SetState(127)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion)|(1<<AddressParserAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAtextStringContext is an interface to support dynamic dispatch.
type IAtextStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtextStringContext differentiates from other interfaces.
	IsAtextStringContext()
}

type AtextStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtextStringContext() *AtextStringContext {
	var p = new(AtextStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_atextString
	return p
}

func (*AtextStringContext) IsAtextStringContext() {}

func NewAtextStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtextStringContext {
	var p = new(AtextStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_atextString

	return p
}

func (s *AtextStringContext) GetParser() antlr.Parser { return s.parser }

func (s *AtextStringContext) AllAtext() []IAtextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtextContext)(nil)).Elem())
	var tst = make([]IAtextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtextContext)
		}
	}

	return tst
}

func (s *AtextStringContext) Atext(i int) IAtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtextContext)
}

func (s *AtextStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtextStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtextStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterAtextString(s)
	}
}

func (s *AtextStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitAtextString(s)
	}
}

func (p *AddressParser) AtextString() (localctx IAtextStringContext) {
	localctx = NewAtextStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AddressParserRULE_atextString)

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
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(129)
				p.Atext()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	p.RuleIndex = AddressParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) AtextString() IAtextStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtextStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtextStringContext)
}

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

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *AddressParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AddressParserRULE_atom)
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
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(134)
			p.Cfws()
		}

	}
	{
		p.SetState(137)
		p.AtextString()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(138)
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
	p.RuleIndex = AddressParserRULE_dotAtomText
	return p
}

func (*DotAtomTextContext) IsDotAtomTextContext() {}

func NewDotAtomTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotAtomTextContext {
	var p = new(DotAtomTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_dotAtomText

	return p
}

func (s *DotAtomTextContext) GetParser() antlr.Parser { return s.parser }

func (s *DotAtomTextContext) AllAtextString() []IAtextStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtextStringContext)(nil)).Elem())
	var tst = make([]IAtextStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtextStringContext)
		}
	}

	return tst
}

func (s *DotAtomTextContext) AtextString(i int) IAtextStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtextStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtextStringContext)
}

func (s *DotAtomTextContext) AllPeriod() []antlr.TerminalNode {
	return s.GetTokens(AddressParserPeriod)
}

func (s *DotAtomTextContext) Period(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, i)
}

func (s *DotAtomTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotAtomTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DotAtomTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterDotAtomText(s)
	}
}

func (s *DotAtomTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitDotAtomText(s)
	}
}

func (p *AddressParser) DotAtomText() (localctx IDotAtomTextContext) {
	localctx = NewDotAtomTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AddressParserRULE_dotAtomText)
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
		p.SetState(141)
		p.AtextString()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserPeriod {
		{
			p.SetState(142)
			p.Match(AddressParserPeriod)
		}
		{
			p.SetState(143)
			p.AtextString()
		}

		p.SetState(148)
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
	p.RuleIndex = AddressParserRULE_dotAtom
	return p
}

func (*DotAtomContext) IsDotAtomContext() {}

func NewDotAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotAtomContext {
	var p = new(DotAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_dotAtom

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterDotAtom(s)
	}
}

func (s *DotAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitDotAtom(s)
	}
}

func (p *AddressParser) DotAtom() (localctx IDotAtomContext) {
	localctx = NewDotAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AddressParserRULE_dotAtom)
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(149)
			p.Cfws()
		}

	}
	{
		p.SetState(152)
		p.DotAtomText()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(153)
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
	p.RuleIndex = AddressParserRULE_qtext
	return p
}

func (*QtextContext) IsQtextContext() {}

func NewQtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QtextContext {
	var p = new(QtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_qtext

	return p
}

func (s *QtextContext) GetParser() antlr.Parser { return s.parser }

func (s *QtextContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressParserExclamation, 0)
}

func (s *QtextContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressParserHash, 0)
}

func (s *QtextContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressParserDollar, 0)
}

func (s *QtextContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressParserPercent, 0)
}

func (s *QtextContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressParserAmpersand, 0)
}

func (s *QtextContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserSQuote, 0)
}

func (s *QtextContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressParserLParens, 0)
}

func (s *QtextContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressParserRParens, 0)
}

func (s *QtextContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressParserAsterisk, 0)
}

func (s *QtextContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressParserPlus, 0)
}

func (s *QtextContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressParserComma, 0)
}

func (s *QtextContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressParserMinus, 0)
}

func (s *QtextContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, 0)
}

func (s *QtextContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressParserSlash, 0)
}

func (s *QtextContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressParserDigit, 0)
}

func (s *QtextContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressParserColon, 0)
}

func (s *QtextContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressParserSemicolon, 0)
}

func (s *QtextContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressParserLess, 0)
}

func (s *QtextContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressParserEqual, 0)
}

func (s *QtextContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressParserGreater, 0)
}

func (s *QtextContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressParserQuestion, 0)
}

func (s *QtextContext) At() antlr.TerminalNode {
	return s.GetToken(AddressParserAt, 0)
}

func (s *QtextContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaUpper, 0)
}

func (s *QtextContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserLBracket, 0)
}

func (s *QtextContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserRBracket, 0)
}

func (s *QtextContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressParserCaret, 0)
}

func (s *QtextContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressParserUnderscore, 0)
}

func (s *QtextContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressParserBacktick, 0)
}

func (s *QtextContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaLower, 0)
}

func (s *QtextContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserLCurly, 0)
}

func (s *QtextContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressParserPipe, 0)
}

func (s *QtextContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserRCurly, 0)
}

func (s *QtextContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressParserTilde, 0)
}

func (s *QtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterQtext(s)
	}
}

func (s *QtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitQtext(s)
	}
}

func (p *AddressParser) Qtext() (localctx IQtextContext) {
	localctx = NewQtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AddressParserRULE_qtext)
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
		p.SetState(156)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater)|(1<<AddressParserQuestion)|(1<<AddressParserAt)|(1<<AddressParserAlphaUpper)|(1<<AddressParserLBracket)|(1<<AddressParserRBracket))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IQuotedContentContext is an interface to support dynamic dispatch.
type IQuotedContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedContentContext differentiates from other interfaces.
	IsQuotedContentContext()
}

type QuotedContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedContentContext() *QuotedContentContext {
	var p = new(QuotedContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_quotedContent
	return p
}

func (*QuotedContentContext) IsQuotedContentContext() {}

func NewQuotedContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedContentContext {
	var p = new(QuotedContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_quotedContent

	return p
}

func (s *QuotedContentContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedContentContext) Qtext() IQtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQtextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQtextContext)
}

func (s *QuotedContentContext) QuotedPair() IQuotedPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedPairContext)
}

func (s *QuotedContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterQuotedContent(s)
	}
}

func (s *QuotedContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitQuotedContent(s)
	}
}

func (p *AddressParser) QuotedContent() (localctx IQuotedContentContext) {
	localctx = NewQuotedContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AddressParserRULE_quotedContent)

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

	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserLParens, AddressParserRParens, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Qtext()
		}

	case AddressParserBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.QuotedPair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuotedValueContext is an interface to support dynamic dispatch.
type IQuotedValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedValueContext differentiates from other interfaces.
	IsQuotedValueContext()
}

type QuotedValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedValueContext() *QuotedValueContext {
	var p = new(QuotedValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_quotedValue
	return p
}

func (*QuotedValueContext) IsQuotedValueContext() {}

func NewQuotedValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedValueContext {
	var p = new(QuotedValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_quotedValue

	return p
}

func (s *QuotedValueContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedValueContext) AllQuotedContent() []IQuotedContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuotedContentContext)(nil)).Elem())
	var tst = make([]IQuotedContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuotedContentContext)
		}
	}

	return tst
}

func (s *QuotedValueContext) QuotedContent(i int) IQuotedContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuotedContentContext)
}

func (s *QuotedValueContext) AllFws() []IFwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFwsContext)(nil)).Elem())
	var tst = make([]IFwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFwsContext)
		}
	}

	return tst
}

func (s *QuotedValueContext) Fws(i int) IFwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFwsContext)
}

func (s *QuotedValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterQuotedValue(s)
	}
}

func (s *QuotedValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitQuotedValue(s)
	}
}

func (p *AddressParser) QuotedValue() (localctx IQuotedValueContext) {
	localctx = NewQuotedValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AddressParserRULE_quotedValue)
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
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(162)
					p.Fws()
				}

			}
			{
				p.SetState(165)
				p.QuotedContent()
			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	p.RuleIndex = AddressParserRULE_quotedString
	return p
}

func (*QuotedStringContext) IsQuotedStringContext() {}

func NewQuotedStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedStringContext {
	var p = new(QuotedStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_quotedString

	return p
}

func (s *QuotedStringContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedStringContext) AllDQuote() []antlr.TerminalNode {
	return s.GetTokens(AddressParserDQuote)
}

func (s *QuotedStringContext) DQuote(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserDQuote, i)
}

func (s *QuotedStringContext) QuotedValue() IQuotedValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedValueContext)
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

func (s *QuotedStringContext) Fws() IFwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFwsContext)(nil)).Elem(), 0)

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterQuotedString(s)
	}
}

func (s *QuotedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitQuotedString(s)
	}
}

func (p *AddressParser) QuotedString() (localctx IQuotedStringContext) {
	localctx = NewQuotedStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AddressParserRULE_quotedString)
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
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(171)
			p.Cfws()
		}

	}
	{
		p.SetState(174)
		p.Match(AddressParserDQuote)
	}
	{
		p.SetState(175)
		p.QuotedValue()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(176)
			p.Fws()
		}

	}
	{
		p.SetState(179)
		p.Match(AddressParserDQuote)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(180)
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
	p.RuleIndex = AddressParserRULE_word
	return p
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_word

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterWord(s)
	}
}

func (s *WordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitWord(s)
	}
}

func (p *AddressParser) Word() (localctx IWordContext) {
	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AddressParserRULE_word)

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

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Atom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.QuotedString()
		}

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
	p.RuleIndex = AddressParserRULE_unstructured
	return p
}

func (*UnstructuredContext) IsUnstructuredContext() {}

func NewUnstructuredContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnstructuredContext {
	var p = new(UnstructuredContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_unstructured

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterUnstructured(s)
	}
}

func (s *UnstructuredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitUnstructured(s)
	}
}

func (p *AddressParser) Unstructured() (localctx IUnstructuredContext) {
	localctx = NewUnstructuredContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AddressParserRULE_unstructured)
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
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(187)
					p.Fws()
				}

			}
			{
				p.SetState(190)
				p.Vchar()
			}

		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserTAB || _la == AddressParserSP {
		{
			p.SetState(196)
			p.Wsp()
		}

		p.SetState(201)
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
	p.RuleIndex = AddressParserRULE_address
	return p
}

func (*AddressContext) IsAddressContext() {}

func NewAddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddressContext {
	var p = new(AddressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_address

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterAddress(s)
	}
}

func (s *AddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitAddress(s)
	}
}

func (p *AddressParser) Address() (localctx IAddressContext) {
	localctx = NewAddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AddressParserRULE_address)

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

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Mailbox()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
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
	p.RuleIndex = AddressParserRULE_mailbox
	return p
}

func (*MailboxContext) IsMailboxContext() {}

func NewMailboxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MailboxContext {
	var p = new(MailboxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_mailbox

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterMailbox(s)
	}
}

func (s *MailboxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitMailbox(s)
	}
}

func (p *AddressParser) Mailbox() (localctx IMailboxContext) {
	localctx = NewMailboxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AddressParserRULE_mailbox)

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

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.NameAddr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
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
	p.RuleIndex = AddressParserRULE_nameAddr
	return p
}

func (*NameAddrContext) IsNameAddrContext() {}

func NewNameAddrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameAddrContext {
	var p = new(NameAddrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_nameAddr

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterNameAddr(s)
	}
}

func (s *NameAddrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitNameAddr(s)
	}
}

func (p *AddressParser) NameAddr() (localctx INameAddrContext) {
	localctx = NewNameAddrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AddressParserRULE_nameAddr)

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
	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(210)
			p.DisplayName()
		}

	}
	{
		p.SetState(213)
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
	p.RuleIndex = AddressParserRULE_angleAddr
	return p
}

func (*AngleAddrContext) IsAngleAddrContext() {}

func NewAngleAddrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AngleAddrContext {
	var p = new(AngleAddrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_angleAddr

	return p
}

func (s *AngleAddrContext) GetParser() antlr.Parser { return s.parser }

func (s *AngleAddrContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressParserLess, 0)
}

func (s *AngleAddrContext) AddrSpec() IAddrSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddrSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddrSpecContext)
}

func (s *AngleAddrContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressParserGreater, 0)
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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterAngleAddr(s)
	}
}

func (s *AngleAddrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitAngleAddr(s)
	}
}

func (p *AddressParser) AngleAddr() (localctx IAngleAddrContext) {
	localctx = NewAngleAddrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AddressParserRULE_angleAddr)
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
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(215)
			p.Cfws()
		}

	}
	{
		p.SetState(218)
		p.Match(AddressParserLess)
	}
	{
		p.SetState(219)
		p.AddrSpec()
	}
	{
		p.SetState(220)
		p.Match(AddressParserGreater)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(221)
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
	p.RuleIndex = AddressParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_group

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
	return s.GetToken(AddressParserColon, 0)
}

func (s *GroupContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressParserSemicolon, 0)
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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *AddressParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AddressParserRULE_group)
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
		p.SetState(224)
		p.DisplayName()
	}
	{
		p.SetState(225)
		p.Match(AddressParserColon)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserQuestion)|(1<<AddressParserAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0) {
		{
			p.SetState(226)
			p.GroupList()
		}

	}
	{
		p.SetState(229)
		p.Match(AddressParserSemicolon)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(230)
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
	p.RuleIndex = AddressParserRULE_displayName
	return p
}

func (*DisplayNameContext) IsDisplayNameContext() {}

func NewDisplayNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisplayNameContext {
	var p = new(DisplayNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_displayName

	return p
}

func (s *DisplayNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DisplayNameContext) AllWord() []IWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordContext)(nil)).Elem())
	var tst = make([]IWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordContext)
		}
	}

	return tst
}

func (s *DisplayNameContext) Word(i int) IWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *DisplayNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisplayNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisplayNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterDisplayName(s)
	}
}

func (s *DisplayNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitDisplayName(s)
	}
}

func (p *AddressParser) DisplayName() (localctx IDisplayNameContext) {
	localctx = NewDisplayNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AddressParserRULE_displayName)

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
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(233)
				p.Word()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
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
	p.RuleIndex = AddressParserRULE_mailboxList
	return p
}

func (*MailboxListContext) IsMailboxListContext() {}

func NewMailboxListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MailboxListContext {
	var p = new(MailboxListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_mailboxList

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
	return s.GetTokens(AddressParserComma)
}

func (s *MailboxListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserComma, i)
}

func (s *MailboxListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MailboxListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MailboxListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterMailboxList(s)
	}
}

func (s *MailboxListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitMailboxList(s)
	}
}

func (p *AddressParser) MailboxList() (localctx IMailboxListContext) {
	localctx = NewMailboxListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AddressParserRULE_mailboxList)
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
		p.SetState(238)
		p.Mailbox()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(239)
			p.Match(AddressParserComma)
		}
		{
			p.SetState(240)
			p.Mailbox()
		}

		p.SetState(245)
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
	p.RuleIndex = AddressParserRULE_addressList
	return p
}

func (*AddressListContext) IsAddressListContext() {}

func NewAddressListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddressListContext {
	var p = new(AddressListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_addressList

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
	return s.GetTokens(AddressParserComma)
}

func (s *AddressListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserComma, i)
}

func (s *AddressListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddressListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddressListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterAddressList(s)
	}
}

func (s *AddressListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitAddressList(s)
	}
}

func (p *AddressParser) AddressList() (localctx IAddressListContext) {
	localctx = NewAddressListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AddressParserRULE_addressList)
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
		p.SetState(246)
		p.Address()
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(247)
			p.Match(AddressParserComma)
		}
		{
			p.SetState(248)
			p.Address()
		}

		p.SetState(253)
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
	p.RuleIndex = AddressParserRULE_groupList
	return p
}

func (*GroupListContext) IsGroupListContext() {}

func NewGroupListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupListContext {
	var p = new(GroupListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_groupList

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterGroupList(s)
	}
}

func (s *GroupListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitGroupList(s)
	}
}

func (p *AddressParser) GroupList() (localctx IGroupListContext) {
	localctx = NewGroupListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AddressParserRULE_groupList)

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

	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.MailboxList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
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
	p.RuleIndex = AddressParserRULE_addrSpec
	return p
}

func (*AddrSpecContext) IsAddrSpecContext() {}

func NewAddrSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddrSpecContext {
	var p = new(AddrSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_addrSpec

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
	return s.GetToken(AddressParserAt, 0)
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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterAddrSpec(s)
	}
}

func (s *AddrSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitAddrSpec(s)
	}
}

func (p *AddressParser) AddrSpec() (localctx IAddrSpecContext) {
	localctx = NewAddrSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AddressParserRULE_addrSpec)

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
		p.SetState(258)
		p.LocalPart()
	}
	{
		p.SetState(259)
		p.Match(AddressParserAt)
	}
	{
		p.SetState(260)
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
	p.RuleIndex = AddressParserRULE_localPart
	return p
}

func (*LocalPartContext) IsLocalPartContext() {}

func NewLocalPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalPartContext {
	var p = new(LocalPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_localPart

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterLocalPart(s)
	}
}

func (s *LocalPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitLocalPart(s)
	}
}

func (p *AddressParser) LocalPart() (localctx ILocalPartContext) {
	localctx = NewLocalPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AddressParserRULE_localPart)

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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.DotAtom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
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
	p.RuleIndex = AddressParserRULE_domain
	return p
}

func (*DomainContext) IsDomainContext() {}

func NewDomainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainContext {
	var p = new(DomainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_domain

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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterDomain(s)
	}
}

func (s *DomainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitDomain(s)
	}
}

func (p *AddressParser) Domain() (localctx IDomainContext) {
	localctx = NewDomainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, AddressParserRULE_domain)

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

	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.DotAtom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
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
	p.RuleIndex = AddressParserRULE_domainLiteral
	return p
}

func (*DomainLiteralContext) IsDomainLiteralContext() {}

func NewDomainLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainLiteralContext {
	var p = new(DomainLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_domainLiteral

	return p
}

func (s *DomainLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainLiteralContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserLBracket, 0)
}

func (s *DomainLiteralContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserRBracket, 0)
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
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterDomainLiteral(s)
	}
}

func (s *DomainLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitDomainLiteral(s)
	}
}

func (p *AddressParser) DomainLiteral() (localctx IDomainLiteralContext) {
	localctx = NewDomainLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AddressParserRULE_domainLiteral)
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
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(270)
			p.Cfws()
		}

	}
	{
		p.SetState(273)
		p.Match(AddressParserLBracket)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(274)
					p.Fws()
				}

			}
			{
				p.SetState(277)
				p.Dtext()
			}

		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(283)
			p.Fws()
		}

	}
	{
		p.SetState(286)
		p.Match(AddressParserRBracket)
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(287)
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
	p.RuleIndex = AddressParserRULE_dtext
	return p
}

func (*DtextContext) IsDtextContext() {}

func NewDtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DtextContext {
	var p = new(DtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_dtext

	return p
}

func (s *DtextContext) GetParser() antlr.Parser { return s.parser }

func (s *DtextContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressParserExclamation, 0)
}

func (s *DtextContext) DQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserDQuote, 0)
}

func (s *DtextContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressParserHash, 0)
}

func (s *DtextContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressParserDollar, 0)
}

func (s *DtextContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressParserPercent, 0)
}

func (s *DtextContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressParserAmpersand, 0)
}

func (s *DtextContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserSQuote, 0)
}

func (s *DtextContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressParserLParens, 0)
}

func (s *DtextContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressParserRParens, 0)
}

func (s *DtextContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressParserAsterisk, 0)
}

func (s *DtextContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressParserPlus, 0)
}

func (s *DtextContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressParserComma, 0)
}

func (s *DtextContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressParserMinus, 0)
}

func (s *DtextContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, 0)
}

func (s *DtextContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressParserSlash, 0)
}

func (s *DtextContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressParserDigit, 0)
}

func (s *DtextContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressParserColon, 0)
}

func (s *DtextContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressParserSemicolon, 0)
}

func (s *DtextContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressParserLess, 0)
}

func (s *DtextContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressParserEqual, 0)
}

func (s *DtextContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressParserGreater, 0)
}

func (s *DtextContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressParserQuestion, 0)
}

func (s *DtextContext) At() antlr.TerminalNode {
	return s.GetToken(AddressParserAt, 0)
}

func (s *DtextContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaUpper, 0)
}

func (s *DtextContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressParserCaret, 0)
}

func (s *DtextContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressParserUnderscore, 0)
}

func (s *DtextContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressParserBacktick, 0)
}

func (s *DtextContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaLower, 0)
}

func (s *DtextContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserLCurly, 0)
}

func (s *DtextContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressParserPipe, 0)
}

func (s *DtextContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserRCurly, 0)
}

func (s *DtextContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressParserTilde, 0)
}

func (s *DtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterDtext(s)
	}
}

func (s *DtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitDtext(s)
	}
}

func (p *AddressParser) Dtext() (localctx IDtextContext) {
	localctx = NewDtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AddressParserRULE_dtext)
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

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater)|(1<<AddressParserQuestion)|(1<<AddressParserAt)|(1<<AddressParserAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICrlfContext is an interface to support dynamic dispatch.
type ICrlfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCrlfContext differentiates from other interfaces.
	IsCrlfContext()
}

type CrlfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCrlfContext() *CrlfContext {
	var p = new(CrlfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_crlf
	return p
}

func (*CrlfContext) IsCrlfContext() {}

func NewCrlfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CrlfContext {
	var p = new(CrlfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_crlf

	return p
}

func (s *CrlfContext) GetParser() antlr.Parser { return s.parser }

func (s *CrlfContext) CR() antlr.TerminalNode {
	return s.GetToken(AddressParserCR, 0)
}

func (s *CrlfContext) LF() antlr.TerminalNode {
	return s.GetToken(AddressParserLF, 0)
}

func (s *CrlfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrlfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CrlfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterCrlf(s)
	}
}

func (s *CrlfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitCrlf(s)
	}
}

func (p *AddressParser) Crlf() (localctx ICrlfContext) {
	localctx = NewCrlfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AddressParserRULE_crlf)

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
		p.SetState(292)
		p.Match(AddressParserCR)
	}
	{
		p.SetState(293)
		p.Match(AddressParserLF)
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
	p.RuleIndex = AddressParserRULE_wsp
	return p
}

func (*WspContext) IsWspContext() {}

func NewWspContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WspContext {
	var p = new(WspContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_wsp

	return p
}

func (s *WspContext) GetParser() antlr.Parser { return s.parser }

func (s *WspContext) SP() antlr.TerminalNode {
	return s.GetToken(AddressParserSP, 0)
}

func (s *WspContext) TAB() antlr.TerminalNode {
	return s.GetToken(AddressParserTAB, 0)
}

func (s *WspContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WspContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WspContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterWsp(s)
	}
}

func (s *WspContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitWsp(s)
	}
}

func (p *AddressParser) Wsp() (localctx IWspContext) {
	localctx = NewWspContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, AddressParserRULE_wsp)
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
		p.SetState(295)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AddressParserTAB || _la == AddressParserSP) {
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
	p.RuleIndex = AddressParserRULE_vchar
	return p
}

func (*VcharContext) IsVcharContext() {}

func NewVcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VcharContext {
	var p = new(VcharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_vchar

	return p
}

func (s *VcharContext) GetParser() antlr.Parser { return s.parser }

func (s *VcharContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressParserExclamation, 0)
}

func (s *VcharContext) DQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserDQuote, 0)
}

func (s *VcharContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressParserHash, 0)
}

func (s *VcharContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressParserDollar, 0)
}

func (s *VcharContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressParserPercent, 0)
}

func (s *VcharContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressParserAmpersand, 0)
}

func (s *VcharContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserSQuote, 0)
}

func (s *VcharContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressParserLParens, 0)
}

func (s *VcharContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressParserRParens, 0)
}

func (s *VcharContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressParserAsterisk, 0)
}

func (s *VcharContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressParserPlus, 0)
}

func (s *VcharContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressParserComma, 0)
}

func (s *VcharContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressParserMinus, 0)
}

func (s *VcharContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, 0)
}

func (s *VcharContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressParserSlash, 0)
}

func (s *VcharContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressParserDigit, 0)
}

func (s *VcharContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressParserColon, 0)
}

func (s *VcharContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressParserSemicolon, 0)
}

func (s *VcharContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressParserLess, 0)
}

func (s *VcharContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressParserEqual, 0)
}

func (s *VcharContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressParserGreater, 0)
}

func (s *VcharContext) Question() antlr.TerminalNode {
	return s.GetToken(AddressParserQuestion, 0)
}

func (s *VcharContext) At() antlr.TerminalNode {
	return s.GetToken(AddressParserAt, 0)
}

func (s *VcharContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaUpper, 0)
}

func (s *VcharContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserLBracket, 0)
}

func (s *VcharContext) Backslash() antlr.TerminalNode {
	return s.GetToken(AddressParserBackslash, 0)
}

func (s *VcharContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserRBracket, 0)
}

func (s *VcharContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressParserCaret, 0)
}

func (s *VcharContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressParserUnderscore, 0)
}

func (s *VcharContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressParserBacktick, 0)
}

func (s *VcharContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaLower, 0)
}

func (s *VcharContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserLCurly, 0)
}

func (s *VcharContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressParserPipe, 0)
}

func (s *VcharContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserRCurly, 0)
}

func (s *VcharContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressParserTilde, 0)
}

func (s *VcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterVchar(s)
	}
}

func (s *VcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitVchar(s)
	}
}

func (p *AddressParser) Vchar() (localctx IVcharContext) {
	localctx = NewVcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, AddressParserRULE_vchar)
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
		p.SetState(297)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater)|(1<<AddressParserQuestion)|(1<<AddressParserAt)|(1<<AddressParserAlphaUpper)|(1<<AddressParserLBracket)|(1<<AddressParserBackslash)|(1<<AddressParserRBracket))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
