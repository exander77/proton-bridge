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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 489,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3, 2, 5,
	2, 99, 10, 2, 3, 3, 3, 3, 3, 3, 3, 4, 7, 4, 105, 10, 4, 12, 4, 14, 4, 108,
	11, 4, 3, 4, 5, 4, 111, 10, 4, 3, 4, 6, 4, 114, 10, 4, 13, 4, 14, 4, 115,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 123, 10, 6, 3, 7, 3, 7, 5, 7, 127,
	10, 7, 3, 7, 7, 7, 130, 10, 7, 12, 7, 14, 7, 133, 11, 7, 3, 7, 5, 7, 136,
	10, 7, 3, 7, 3, 7, 3, 8, 5, 8, 141, 10, 8, 3, 8, 6, 8, 144, 10, 8, 13,
	8, 14, 8, 145, 3, 8, 5, 8, 149, 10, 8, 3, 8, 5, 8, 152, 10, 8, 3, 9, 3,
	9, 3, 10, 6, 10, 157, 10, 10, 13, 10, 14, 10, 158, 3, 11, 6, 11, 162, 10,
	11, 13, 11, 14, 11, 163, 3, 11, 3, 11, 6, 11, 168, 10, 11, 13, 11, 14,
	11, 169, 7, 11, 172, 10, 11, 12, 11, 14, 11, 175, 11, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 5, 13, 181, 10, 13, 3, 14, 5, 14, 184, 10, 14, 3, 14, 7,
	14, 187, 10, 14, 12, 14, 14, 14, 190, 11, 14, 3, 15, 3, 15, 3, 15, 5, 15,
	195, 10, 15, 3, 15, 3, 15, 3, 16, 5, 16, 200, 10, 16, 3, 16, 3, 16, 5,
	16, 204, 10, 16, 3, 16, 5, 16, 207, 10, 16, 3, 16, 3, 16, 5, 16, 211, 10,
	16, 3, 16, 5, 16, 214, 10, 16, 3, 16, 3, 16, 5, 16, 218, 10, 16, 5, 16,
	220, 10, 16, 3, 17, 5, 17, 223, 10, 17, 3, 17, 7, 17, 226, 10, 17, 12,
	17, 14, 17, 229, 11, 17, 3, 17, 7, 17, 232, 10, 17, 12, 17, 14, 17, 235,
	11, 17, 3, 18, 3, 18, 5, 18, 239, 10, 18, 3, 19, 3, 19, 5, 19, 243, 10,
	19, 3, 20, 5, 20, 246, 10, 20, 3, 20, 3, 20, 3, 21, 5, 21, 251, 10, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 257, 10, 21, 3, 21, 5, 21, 260, 10,
	21, 3, 22, 3, 22, 3, 22, 5, 22, 265, 10, 22, 3, 22, 3, 22, 5, 22, 269,
	10, 22, 3, 23, 6, 23, 272, 10, 23, 13, 23, 14, 23, 273, 3, 23, 5, 23, 277,
	10, 23, 3, 24, 3, 24, 3, 24, 7, 24, 282, 10, 24, 12, 24, 14, 24, 285, 11,
	24, 3, 24, 5, 24, 288, 10, 24, 3, 25, 3, 25, 3, 25, 7, 25, 293, 10, 25,
	12, 25, 14, 25, 296, 11, 25, 3, 25, 5, 25, 299, 10, 25, 3, 26, 3, 26, 3,
	26, 5, 26, 304, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 5, 28, 311,
	10, 28, 3, 28, 3, 28, 5, 28, 315, 10, 28, 3, 28, 5, 28, 318, 10, 28, 3,
	28, 3, 28, 5, 28, 322, 10, 28, 5, 28, 324, 10, 28, 3, 29, 5, 29, 327, 10,
	29, 3, 29, 3, 29, 5, 29, 331, 10, 29, 3, 29, 5, 29, 334, 10, 29, 3, 29,
	3, 29, 5, 29, 338, 10, 29, 5, 29, 340, 10, 29, 3, 30, 3, 30, 5, 30, 344,
	10, 30, 3, 30, 7, 30, 347, 10, 30, 12, 30, 14, 30, 350, 11, 30, 3, 30,
	5, 30, 353, 10, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 7, 32, 363, 10, 32, 12, 32, 14, 32, 366, 11, 32, 3, 33, 5, 33, 369,
	10, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 376, 10, 33, 3, 34, 3,
	34, 3, 34, 3, 35, 3, 35, 7, 35, 383, 10, 35, 12, 35, 14, 35, 386, 11, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 392, 10, 35, 3, 35, 3, 35, 5, 35, 396,
	10, 35, 7, 35, 398, 10, 35, 12, 35, 14, 35, 401, 11, 35, 3, 36, 5, 36,
	404, 10, 36, 3, 36, 7, 36, 407, 10, 36, 12, 36, 14, 36, 410, 11, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 5, 36, 416, 10, 36, 7, 36, 418, 10, 36, 12, 36,
	14, 36, 421, 11, 36, 3, 37, 5, 37, 424, 10, 37, 3, 37, 7, 37, 427, 10,
	37, 12, 37, 14, 37, 430, 11, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 436,
	10, 37, 7, 37, 438, 10, 37, 12, 37, 14, 37, 441, 11, 37, 3, 38, 5, 38,
	444, 10, 38, 3, 38, 6, 38, 447, 10, 38, 13, 38, 14, 38, 448, 3, 38, 5,
	38, 452, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 6, 42, 469, 10, 42, 13,
	42, 14, 42, 470, 3, 43, 3, 43, 3, 44, 6, 44, 476, 10, 44, 13, 44, 14, 44,
	477, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3,
	48, 2, 2, 49, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 2, 10, 5, 2, 7, 13,
	16, 31, 33, 41, 11, 2, 7, 7, 9, 13, 16, 17, 19, 19, 21, 22, 26, 26, 28,
	28, 30, 30, 34, 41, 5, 2, 7, 7, 9, 31, 33, 41, 4, 2, 7, 30, 34, 41, 10,
	2, 7, 7, 9, 13, 16, 17, 19, 19, 22, 22, 30, 30, 32, 32, 34, 41, 4, 2, 7,
	27, 29, 41, 4, 2, 3, 3, 6, 6, 3, 2, 7, 41, 2, 527, 2, 98, 3, 2, 2, 2, 4,
	100, 3, 2, 2, 2, 6, 110, 3, 2, 2, 2, 8, 117, 3, 2, 2, 2, 10, 122, 3, 2,
	2, 2, 12, 124, 3, 2, 2, 2, 14, 151, 3, 2, 2, 2, 16, 153, 3, 2, 2, 2, 18,
	156, 3, 2, 2, 2, 20, 161, 3, 2, 2, 2, 22, 176, 3, 2, 2, 2, 24, 180, 3,
	2, 2, 2, 26, 188, 3, 2, 2, 2, 28, 191, 3, 2, 2, 2, 30, 219, 3, 2, 2, 2,
	32, 227, 3, 2, 2, 2, 34, 238, 3, 2, 2, 2, 36, 242, 3, 2, 2, 2, 38, 245,
	3, 2, 2, 2, 40, 259, 3, 2, 2, 2, 42, 261, 3, 2, 2, 2, 44, 276, 3, 2, 2,
	2, 46, 287, 3, 2, 2, 2, 48, 298, 3, 2, 2, 2, 50, 303, 3, 2, 2, 2, 52, 305,
	3, 2, 2, 2, 54, 323, 3, 2, 2, 2, 56, 339, 3, 2, 2, 2, 58, 341, 3, 2, 2,
	2, 60, 356, 3, 2, 2, 2, 62, 358, 3, 2, 2, 2, 64, 368, 3, 2, 2, 2, 66, 377,
	3, 2, 2, 2, 68, 384, 3, 2, 2, 2, 70, 408, 3, 2, 2, 2, 72, 428, 3, 2, 2,
	2, 74, 446, 3, 2, 2, 2, 76, 453, 3, 2, 2, 2, 78, 463, 3, 2, 2, 2, 80, 465,
	3, 2, 2, 2, 82, 468, 3, 2, 2, 2, 84, 472, 3, 2, 2, 2, 86, 475, 3, 2, 2,
	2, 88, 479, 3, 2, 2, 2, 90, 481, 3, 2, 2, 2, 92, 484, 3, 2, 2, 2, 94, 486,
	3, 2, 2, 2, 96, 99, 5, 94, 48, 2, 97, 99, 5, 92, 47, 2, 98, 96, 3, 2, 2,
	2, 98, 97, 3, 2, 2, 2, 99, 3, 3, 2, 2, 2, 100, 101, 7, 32, 2, 2, 101, 102,
	5, 2, 2, 2, 102, 5, 3, 2, 2, 2, 103, 105, 5, 92, 47, 2, 104, 103, 3, 2,
	2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2,
	107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 111, 5, 90, 46, 2, 110,
	106, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 113, 3, 2, 2, 2, 112, 114,
	5, 92, 47, 2, 113, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 113, 3,
	2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 7, 3, 2, 2, 2, 117, 118, 9, 2, 2, 2,
	118, 9, 3, 2, 2, 2, 119, 123, 5, 8, 5, 2, 120, 123, 5, 4, 3, 2, 121, 123,
	5, 12, 7, 2, 122, 119, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 121, 3, 2,
	2, 2, 123, 11, 3, 2, 2, 2, 124, 131, 7, 14, 2, 2, 125, 127, 5, 6, 4, 2,
	126, 125, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128,
	130, 5, 10, 6, 2, 129, 126, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 129,
	3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2,
	2, 2, 134, 136, 5, 6, 4, 2, 135, 134, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2,
	136, 137, 3, 2, 2, 2, 137, 138, 7, 15, 2, 2, 138, 13, 3, 2, 2, 2, 139,
	141, 5, 6, 4, 2, 140, 139, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142,
	3, 2, 2, 2, 142, 144, 5, 12, 7, 2, 143, 140, 3, 2, 2, 2, 144, 145, 3, 2,
	2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 148, 3, 2, 2, 2,
	147, 149, 5, 6, 4, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	152, 3, 2, 2, 2, 150, 152, 5, 6, 4, 2, 151, 143, 3, 2, 2, 2, 151, 150,
	3, 2, 2, 2, 152, 15, 3, 2, 2, 2, 153, 154, 9, 3, 2, 2, 154, 17, 3, 2, 2,
	2, 155, 157, 5, 16, 9, 2, 156, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158,
	156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 19, 3, 2, 2, 2, 160, 162, 5,
	16, 9, 2, 161, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 161, 3, 2, 2,
	2, 163, 164, 3, 2, 2, 2, 164, 173, 3, 2, 2, 2, 165, 167, 7, 20, 2, 2, 166,
	168, 5, 16, 9, 2, 167, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 167,
	3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 172, 3, 2, 2, 2, 171, 165, 3, 2,
	2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2,
	174, 21, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176, 177, 9, 4, 2, 2, 177, 23,
	3, 2, 2, 2, 178, 181, 5, 22, 12, 2, 179, 181, 5, 4, 3, 2, 180, 178, 3,
	2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 25, 3, 2, 2, 2, 182, 184, 5, 6, 4,
	2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185,
	187, 5, 24, 13, 2, 186, 183, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188, 186,
	3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 27, 3, 2, 2, 2, 190, 188, 3, 2,
	2, 2, 191, 192, 7, 8, 2, 2, 192, 194, 5, 26, 14, 2, 193, 195, 5, 6, 4,
	2, 194, 193, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196,
	197, 7, 8, 2, 2, 197, 29, 3, 2, 2, 2, 198, 200, 5, 14, 8, 2, 199, 198,
	3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 203, 5, 76,
	39, 2, 202, 204, 5, 14, 8, 2, 203, 202, 3, 2, 2, 2, 203, 204, 3, 2, 2,
	2, 204, 220, 3, 2, 2, 2, 205, 207, 5, 14, 8, 2, 206, 205, 3, 2, 2, 2, 206,
	207, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 5, 18, 10, 2, 209, 211,
	5, 14, 8, 2, 210, 209, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 220, 3, 2,
	2, 2, 212, 214, 5, 14, 8, 2, 213, 212, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2,
	214, 215, 3, 2, 2, 2, 215, 217, 5, 28, 15, 2, 216, 218, 5, 14, 8, 2, 217,
	216, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 220, 3, 2, 2, 2, 219, 199,
	3, 2, 2, 2, 219, 206, 3, 2, 2, 2, 219, 213, 3, 2, 2, 2, 220, 31, 3, 2,
	2, 2, 221, 223, 5, 6, 4, 2, 222, 221, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2,
	223, 224, 3, 2, 2, 2, 224, 226, 5, 94, 48, 2, 225, 222, 3, 2, 2, 2, 226,
	229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 233,
	3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 232, 5, 92, 47, 2, 231, 230, 3,
	2, 2, 2, 232, 235, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2,
	2, 234, 33, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 236, 239, 5, 36, 19, 2, 237,
	239, 5, 42, 22, 2, 238, 236, 3, 2, 2, 2, 238, 237, 3, 2, 2, 2, 239, 35,
	3, 2, 2, 2, 240, 243, 5, 38, 20, 2, 241, 243, 5, 52, 27, 2, 242, 240, 3,
	2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 37, 3, 2, 2, 2, 244, 246, 5, 44, 23,
	2, 245, 244, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247,
	248, 5, 40, 21, 2, 248, 39, 3, 2, 2, 2, 249, 251, 5, 14, 8, 2, 250, 249,
	3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 7, 25,
	2, 2, 253, 254, 5, 52, 27, 2, 254, 256, 7, 27, 2, 2, 255, 257, 5, 14, 8,
	2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258,
	260, 5, 64, 33, 2, 259, 250, 3, 2, 2, 2, 259, 258, 3, 2, 2, 2, 260, 41,
	3, 2, 2, 2, 261, 262, 5, 44, 23, 2, 262, 264, 7, 23, 2, 2, 263, 265, 5,
	50, 26, 2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 266, 3, 2,
	2, 2, 266, 268, 7, 24, 2, 2, 267, 269, 5, 14, 8, 2, 268, 267, 3, 2, 2,
	2, 268, 269, 3, 2, 2, 2, 269, 43, 3, 2, 2, 2, 270, 272, 5, 30, 16, 2, 271,
	270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 271, 3, 2, 2, 2, 273, 274,
	3, 2, 2, 2, 274, 277, 3, 2, 2, 2, 275, 277, 5, 62, 32, 2, 276, 271, 3,
	2, 2, 2, 276, 275, 3, 2, 2, 2, 277, 45, 3, 2, 2, 2, 278, 283, 5, 36, 19,
	2, 279, 280, 7, 18, 2, 2, 280, 282, 5, 36, 19, 2, 281, 279, 3, 2, 2, 2,
	282, 285, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284,
	288, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 286, 288, 5, 70, 36, 2, 287, 278,
	3, 2, 2, 2, 287, 286, 3, 2, 2, 2, 288, 47, 3, 2, 2, 2, 289, 294, 5, 34,
	18, 2, 290, 291, 7, 18, 2, 2, 291, 293, 5, 34, 18, 2, 292, 290, 3, 2, 2,
	2, 293, 296, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295,
	299, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 297, 299, 5, 72, 37, 2, 298, 289,
	3, 2, 2, 2, 298, 297, 3, 2, 2, 2, 299, 49, 3, 2, 2, 2, 300, 304, 5, 46,
	24, 2, 301, 304, 5, 14, 8, 2, 302, 304, 5, 74, 38, 2, 303, 300, 3, 2, 2,
	2, 303, 301, 3, 2, 2, 2, 303, 302, 3, 2, 2, 2, 304, 51, 3, 2, 2, 2, 305,
	306, 5, 54, 28, 2, 306, 307, 7, 29, 2, 2, 307, 308, 5, 56, 29, 2, 308,
	53, 3, 2, 2, 2, 309, 311, 5, 14, 8, 2, 310, 309, 3, 2, 2, 2, 310, 311,
	3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312, 314, 5, 20, 11, 2, 313, 315, 5,
	14, 8, 2, 314, 313, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 324, 3, 2, 2,
	2, 316, 318, 5, 14, 8, 2, 317, 316, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318,
	319, 3, 2, 2, 2, 319, 321, 5, 28, 15, 2, 320, 322, 5, 14, 8, 2, 321, 320,
	3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 324, 3, 2, 2, 2, 323, 310, 3, 2,
	2, 2, 323, 317, 3, 2, 2, 2, 324, 55, 3, 2, 2, 2, 325, 327, 5, 14, 8, 2,
	326, 325, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328,
	330, 5, 20, 11, 2, 329, 331, 5, 14, 8, 2, 330, 329, 3, 2, 2, 2, 330, 331,
	3, 2, 2, 2, 331, 340, 3, 2, 2, 2, 332, 334, 5, 14, 8, 2, 333, 332, 3, 2,
	2, 2, 333, 334, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 337, 5, 58, 30,
	2, 336, 338, 5, 14, 8, 2, 337, 336, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338,
	340, 3, 2, 2, 2, 339, 326, 3, 2, 2, 2, 339, 333, 3, 2, 2, 2, 340, 57, 3,
	2, 2, 2, 341, 348, 7, 31, 2, 2, 342, 344, 5, 6, 4, 2, 343, 342, 3, 2, 2,
	2, 343, 344, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 347, 5, 60, 31, 2,
	346, 343, 3, 2, 2, 2, 347, 350, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348,
	349, 3, 2, 2, 2, 349, 352, 3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 351, 353,
	5, 6, 4, 2, 352, 351, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 354, 3, 2,
	2, 2, 354, 355, 7, 33, 2, 2, 355, 59, 3, 2, 2, 2, 356, 357, 9, 5, 2, 2,
	357, 61, 3, 2, 2, 2, 358, 364, 5, 30, 16, 2, 359, 363, 5, 30, 16, 2, 360,
	363, 7, 20, 2, 2, 361, 363, 5, 14, 8, 2, 362, 359, 3, 2, 2, 2, 362, 360,
	3, 2, 2, 2, 362, 361, 3, 2, 2, 2, 363, 366, 3, 2, 2, 2, 364, 362, 3, 2,
	2, 2, 364, 365, 3, 2, 2, 2, 365, 63, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2,
	367, 369, 5, 14, 8, 2, 368, 367, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369,
	370, 3, 2, 2, 2, 370, 371, 7, 25, 2, 2, 371, 372, 5, 66, 34, 2, 372, 373,
	5, 52, 27, 2, 373, 375, 7, 27, 2, 2, 374, 376, 5, 14, 8, 2, 375, 374, 3,
	2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 65, 3, 2, 2, 2, 377, 378, 5, 68, 35,
	2, 378, 379, 7, 23, 2, 2, 379, 67, 3, 2, 2, 2, 380, 383, 5, 14, 8, 2, 381,
	383, 7, 18, 2, 2, 382, 380, 3, 2, 2, 2, 382, 381, 3, 2, 2, 2, 383, 386,
	3, 2, 2, 2, 384, 382, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 387, 3, 2,
	2, 2, 386, 384, 3, 2, 2, 2, 387, 388, 7, 29, 2, 2, 388, 399, 5, 56, 29,
	2, 389, 391, 7, 18, 2, 2, 390, 392, 5, 14, 8, 2, 391, 390, 3, 2, 2, 2,
	391, 392, 3, 2, 2, 2, 392, 395, 3, 2, 2, 2, 393, 394, 7, 29, 2, 2, 394,
	396, 5, 56, 29, 2, 395, 393, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 398,
	3, 2, 2, 2, 397, 389, 3, 2, 2, 2, 398, 401, 3, 2, 2, 2, 399, 397, 3, 2,
	2, 2, 399, 400, 3, 2, 2, 2, 400, 69, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2,
	402, 404, 5, 14, 8, 2, 403, 402, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404,
	405, 3, 2, 2, 2, 405, 407, 7, 18, 2, 2, 406, 403, 3, 2, 2, 2, 407, 410,
	3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 411, 3, 2,
	2, 2, 410, 408, 3, 2, 2, 2, 411, 419, 5, 36, 19, 2, 412, 415, 7, 18, 2,
	2, 413, 416, 5, 36, 19, 2, 414, 416, 5, 14, 8, 2, 415, 413, 3, 2, 2, 2,
	415, 414, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 418, 3, 2, 2, 2, 417,
	412, 3, 2, 2, 2, 418, 421, 3, 2, 2, 2, 419, 417, 3, 2, 2, 2, 419, 420,
	3, 2, 2, 2, 420, 71, 3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 422, 424, 5, 14,
	8, 2, 423, 422, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2,
	425, 427, 7, 18, 2, 2, 426, 423, 3, 2, 2, 2, 427, 430, 3, 2, 2, 2, 428,
	426, 3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 429, 431, 3, 2, 2, 2, 430, 428,
	3, 2, 2, 2, 431, 439, 5, 34, 18, 2, 432, 435, 7, 18, 2, 2, 433, 436, 5,
	34, 18, 2, 434, 436, 5, 14, 8, 2, 435, 433, 3, 2, 2, 2, 435, 434, 3, 2,
	2, 2, 435, 436, 3, 2, 2, 2, 436, 438, 3, 2, 2, 2, 437, 432, 3, 2, 2, 2,
	438, 441, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440,
	73, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 442, 444, 5, 14, 8, 2, 443, 442,
	3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 447, 7, 18,
	2, 2, 446, 443, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 446, 3, 2, 2, 2,
	448, 449, 3, 2, 2, 2, 449, 451, 3, 2, 2, 2, 450, 452, 5, 14, 8, 2, 451,
	450, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 452, 75, 3, 2, 2, 2, 453, 454, 7,
	26, 2, 2, 454, 455, 7, 28, 2, 2, 455, 456, 5, 78, 40, 2, 456, 457, 7, 28,
	2, 2, 457, 458, 5, 80, 41, 2, 458, 459, 7, 28, 2, 2, 459, 460, 5, 86, 44,
	2, 460, 461, 7, 28, 2, 2, 461, 462, 7, 26, 2, 2, 462, 77, 3, 2, 2, 2, 463,
	464, 5, 82, 42, 2, 464, 79, 3, 2, 2, 2, 465, 466, 5, 82, 42, 2, 466, 81,
	3, 2, 2, 2, 467, 469, 5, 84, 43, 2, 468, 467, 3, 2, 2, 2, 469, 470, 3,
	2, 2, 2, 470, 468, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 83, 3, 2, 2,
	2, 472, 473, 9, 6, 2, 2, 473, 85, 3, 2, 2, 2, 474, 476, 5, 88, 45, 2, 475,
	474, 3, 2, 2, 2, 476, 477, 3, 2, 2, 2, 477, 475, 3, 2, 2, 2, 477, 478,
	3, 2, 2, 2, 478, 87, 3, 2, 2, 2, 479, 480, 9, 7, 2, 2, 480, 89, 3, 2, 2,
	2, 481, 482, 7, 5, 2, 2, 482, 483, 7, 4, 2, 2, 483, 91, 3, 2, 2, 2, 484,
	485, 9, 8, 2, 2, 485, 93, 3, 2, 2, 2, 486, 487, 9, 9, 2, 2, 487, 95, 3,
	2, 2, 2, 82, 98, 106, 110, 115, 122, 126, 131, 135, 140, 145, 148, 151,
	158, 163, 169, 173, 180, 183, 188, 194, 199, 203, 206, 210, 213, 217, 219,
	222, 227, 233, 238, 242, 245, 250, 256, 259, 264, 268, 273, 276, 283, 287,
	294, 298, 303, 310, 314, 317, 321, 323, 326, 330, 333, 337, 339, 343, 348,
	352, 362, 364, 368, 375, 382, 384, 391, 395, 399, 403, 408, 415, 419, 423,
	428, 435, 439, 443, 448, 451, 470, 477,
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
	"atext", "atom", "dotAtom", "qtext", "quotedContent", "quotedValue", "quotedString",
	"word", "unstructured", "address", "mailbox", "nameAddr", "angleAddr",
	"group", "displayName", "mailboxList", "addressList", "groupList", "addrSpec",
	"localPart", "domain", "domainLiteral", "dtext", "obsPhrase", "obsAngleAddr",
	"obsRoute", "obsDomainList", "obsMboxList", "obsAddrList", "obsGroupList",
	"encodedWord", "charset", "encoding", "token", "tokenChar", "encodedText",
	"encodedChar", "crlf", "wsp", "vchar",
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
	AddressParserRULE_atom          = 8
	AddressParserRULE_dotAtom       = 9
	AddressParserRULE_qtext         = 10
	AddressParserRULE_quotedContent = 11
	AddressParserRULE_quotedValue   = 12
	AddressParserRULE_quotedString  = 13
	AddressParserRULE_word          = 14
	AddressParserRULE_unstructured  = 15
	AddressParserRULE_address       = 16
	AddressParserRULE_mailbox       = 17
	AddressParserRULE_nameAddr      = 18
	AddressParserRULE_angleAddr     = 19
	AddressParserRULE_group         = 20
	AddressParserRULE_displayName   = 21
	AddressParserRULE_mailboxList   = 22
	AddressParserRULE_addressList   = 23
	AddressParserRULE_groupList     = 24
	AddressParserRULE_addrSpec      = 25
	AddressParserRULE_localPart     = 26
	AddressParserRULE_domain        = 27
	AddressParserRULE_domainLiteral = 28
	AddressParserRULE_dtext         = 29
	AddressParserRULE_obsPhrase     = 30
	AddressParserRULE_obsAngleAddr  = 31
	AddressParserRULE_obsRoute      = 32
	AddressParserRULE_obsDomainList = 33
	AddressParserRULE_obsMboxList   = 34
	AddressParserRULE_obsAddrList   = 35
	AddressParserRULE_obsGroupList  = 36
	AddressParserRULE_encodedWord   = 37
	AddressParserRULE_charset       = 38
	AddressParserRULE_encoding      = 39
	AddressParserRULE_token         = 40
	AddressParserRULE_tokenChar     = 41
	AddressParserRULE_encodedText   = 42
	AddressParserRULE_encodedChar   = 43
	AddressParserRULE_crlf          = 44
	AddressParserRULE_wsp           = 45
	AddressParserRULE_vchar         = 46
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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation, AddressParserDQuote, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserLParens, AddressParserRParens, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserBackslash, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Vchar()
		}

	case AddressParserTAB, AddressParserSP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
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
		p.SetState(98)
		p.Match(AddressParserBackslash)
	}
	{
		p.SetState(99)
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
	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AddressParserTAB || _la == AddressParserSP {
			{
				p.SetState(101)
				p.Wsp()
			}

			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
			p.Crlf()
		}

	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(110)
				p.Wsp()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(113)
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
		p.SetState(115)
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

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation, AddressParserDQuote, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Ctext()
		}

	case AddressParserBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.QuotedPair()
		}

	case AddressParserLParens:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
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
		p.SetState(122)
		p.Match(AddressParserLParens)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(123)
					p.Fws()
				}

			}
			{
				p.SetState(126)
				p.Ccontent()
			}

		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(132)
			p.Fws()
		}

	}
	{
		p.SetState(135)
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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(138)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
					{
						p.SetState(137)
						p.Fws()
					}

				}
				{
					p.SetState(140)
					p.Comment()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(145)
				p.Fws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
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
		p.SetState(151)
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
	p.EnterRule(localctx, 16, AddressParserRULE_atom)

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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(153)
				p.Atext()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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

func (s *DotAtomContext) AllAtext() []IAtextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtextContext)(nil)).Elem())
	var tst = make([]IAtextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtextContext)
		}
	}

	return tst
}

func (s *DotAtomContext) Atext(i int) IAtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtextContext)
}

func (s *DotAtomContext) AllPeriod() []antlr.TerminalNode {
	return s.GetTokens(AddressParserPeriod)
}

func (s *DotAtomContext) Period(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, i)
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
	p.EnterRule(localctx, 18, AddressParserRULE_dotAtom)
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
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion)|(1<<AddressParserAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0) {
		{
			p.SetState(158)
			p.Atext()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserPeriod {
		{
			p.SetState(163)
			p.Match(AddressParserPeriod)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion)|(1<<AddressParserAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0) {
			{
				p.SetState(164)
				p.Atext()
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 20, AddressParserRULE_qtext)
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
		p.SetState(174)
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
	p.EnterRule(localctx, 22, AddressParserRULE_quotedContent)

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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserLParens, AddressParserRParens, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Qtext()
		}

	case AddressParserBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
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
	p.EnterRule(localctx, 24, AddressParserRULE_quotedValue)
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
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(180)
					p.Fws()
				}

			}
			{
				p.SetState(183)
				p.QuotedContent()
			}

		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, AddressParserRULE_quotedString)
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
		p.SetState(189)
		p.Match(AddressParserDQuote)
	}
	{
		p.SetState(190)
		p.QuotedValue()
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(191)
			p.Fws()
		}

	}
	{
		p.SetState(194)
		p.Match(AddressParserDQuote)
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

func (s *WordContext) EncodedWord() IEncodedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncodedWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEncodedWordContext)
}

func (s *WordContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *WordContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

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
	p.EnterRule(localctx, 28, AddressParserRULE_word)
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

	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(196)
				p.Cfws()
			}

		}
		{
			p.SetState(199)
			p.EncodedWord()
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(200)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(203)
				p.Cfws()
			}

		}
		{
			p.SetState(206)
			p.Atom()
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(207)
				p.Cfws()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(210)
				p.Cfws()
			}

		}
		{
			p.SetState(213)
			p.QuotedString()
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(214)
				p.Cfws()
			}

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
	p.EnterRule(localctx, 30, AddressParserRULE_unstructured)
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
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(219)
					p.Fws()
				}

			}
			{
				p.SetState(222)
				p.Vchar()
			}

		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserTAB || _la == AddressParserSP {
		{
			p.SetState(228)
			p.Wsp()
		}

		p.SetState(233)
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
	p.EnterRule(localctx, 32, AddressParserRULE_address)

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

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Mailbox()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
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
	p.EnterRule(localctx, 34, AddressParserRULE_mailbox)

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

	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.NameAddr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
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
	p.EnterRule(localctx, 36, AddressParserRULE_nameAddr)

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
	p.SetState(243)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(242)
			p.DisplayName()
		}

	}
	{
		p.SetState(245)
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

func (s *AngleAddrContext) ObsAngleAddr() IObsAngleAddrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsAngleAddrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsAngleAddrContext)
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
	p.EnterRule(localctx, 38, AddressParserRULE_angleAddr)
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

	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(247)
				p.Cfws()
			}

		}
		{
			p.SetState(250)
			p.Match(AddressParserLess)
		}
		{
			p.SetState(251)
			p.AddrSpec()
		}
		{
			p.SetState(252)
			p.Match(AddressParserGreater)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(253)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.ObsAngleAddr()
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
	p.EnterRule(localctx, 40, AddressParserRULE_group)
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
		p.SetState(259)
		p.DisplayName()
	}
	{
		p.SetState(260)
		p.Match(AddressParserColon)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserQuestion)|(1<<AddressParserAlphaUpper))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0) {
		{
			p.SetState(261)
			p.GroupList()
		}

	}
	{
		p.SetState(264)
		p.Match(AddressParserSemicolon)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(265)
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

func (s *DisplayNameContext) ObsPhrase() IObsPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsPhraseContext)
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
	p.EnterRule(localctx, 42, AddressParserRULE_displayName)

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

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(268)
					p.Word()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.ObsPhrase()
		}

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

func (s *MailboxListContext) ObsMboxList() IObsMboxListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsMboxListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsMboxListContext)
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
	p.EnterRule(localctx, 44, AddressParserRULE_mailboxList)
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

	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.Mailbox()
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AddressParserComma {
			{
				p.SetState(277)
				p.Match(AddressParserComma)
			}
			{
				p.SetState(278)
				p.Mailbox()
			}

			p.SetState(283)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(284)
			p.ObsMboxList()
		}

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

func (s *AddressListContext) ObsAddrList() IObsAddrListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsAddrListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsAddrListContext)
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
	p.EnterRule(localctx, 46, AddressParserRULE_addressList)
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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Address()
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AddressParserComma {
			{
				p.SetState(288)
				p.Match(AddressParserComma)
			}
			{
				p.SetState(289)
				p.Address()
			}

			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.ObsAddrList()
		}

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

func (s *GroupListContext) ObsGroupList() IObsGroupListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsGroupListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsGroupListContext)
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
	p.EnterRule(localctx, 48, AddressParserRULE_groupList)

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

	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.MailboxList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.Cfws()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(300)
			p.ObsGroupList()
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
	p.EnterRule(localctx, 50, AddressParserRULE_addrSpec)

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
		p.SetState(303)
		p.LocalPart()
	}
	{
		p.SetState(304)
		p.Match(AddressParserAt)
	}
	{
		p.SetState(305)
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

func (s *LocalPartContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *LocalPartContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
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
	p.EnterRule(localctx, 52, AddressParserRULE_localPart)
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

	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(307)
				p.Cfws()
			}

		}
		{
			p.SetState(310)
			p.DotAtom()
		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(311)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(314)
				p.Cfws()
			}

		}
		{
			p.SetState(317)
			p.QuotedString()
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(318)
				p.Cfws()
			}

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

func (s *DomainContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *DomainContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
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
	p.EnterRule(localctx, 54, AddressParserRULE_domain)
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

	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(323)
				p.Cfws()
			}

		}
		{
			p.SetState(326)
			p.DotAtom()
		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(327)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(330)
				p.Cfws()
			}

		}
		{
			p.SetState(333)
			p.DomainLiteral()
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(334)
				p.Cfws()
			}

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
	p.EnterRule(localctx, 56, AddressParserRULE_domainLiteral)
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
		p.SetState(339)
		p.Match(AddressParserLBracket)
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(341)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(340)
					p.Fws()
				}

			}
			{
				p.SetState(343)
				p.Dtext()
			}

		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(349)
			p.Fws()
		}

	}
	{
		p.SetState(352)
		p.Match(AddressParserRBracket)
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
	p.EnterRule(localctx, 58, AddressParserRULE_dtext)
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
		p.SetState(354)
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

// IObsPhraseContext is an interface to support dynamic dispatch.
type IObsPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsPhraseContext differentiates from other interfaces.
	IsObsPhraseContext()
}

type ObsPhraseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsPhraseContext() *ObsPhraseContext {
	var p = new(ObsPhraseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsPhrase
	return p
}

func (*ObsPhraseContext) IsObsPhraseContext() {}

func NewObsPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsPhraseContext {
	var p = new(ObsPhraseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsPhrase

	return p
}

func (s *ObsPhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsPhraseContext) AllWord() []IWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordContext)(nil)).Elem())
	var tst = make([]IWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordContext)
		}
	}

	return tst
}

func (s *ObsPhraseContext) Word(i int) IWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *ObsPhraseContext) AllPeriod() []antlr.TerminalNode {
	return s.GetTokens(AddressParserPeriod)
}

func (s *ObsPhraseContext) Period(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, i)
}

func (s *ObsPhraseContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *ObsPhraseContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *ObsPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsPhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsPhraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsPhrase(s)
	}
}

func (s *ObsPhraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsPhrase(s)
	}
}

func (p *AddressParser) ObsPhrase() (localctx IObsPhraseContext) {
	localctx = NewObsPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AddressParserRULE_obsPhrase)

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
		p.SetState(356)
		p.Word()
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(357)
					p.Word()
				}

			case 2:
				{
					p.SetState(358)
					p.Match(AddressParserPeriod)
				}

			case 3:
				{
					p.SetState(359)
					p.Cfws()
				}

			}

		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())
	}

	return localctx
}

// IObsAngleAddrContext is an interface to support dynamic dispatch.
type IObsAngleAddrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsAngleAddrContext differentiates from other interfaces.
	IsObsAngleAddrContext()
}

type ObsAngleAddrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsAngleAddrContext() *ObsAngleAddrContext {
	var p = new(ObsAngleAddrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsAngleAddr
	return p
}

func (*ObsAngleAddrContext) IsObsAngleAddrContext() {}

func NewObsAngleAddrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsAngleAddrContext {
	var p = new(ObsAngleAddrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsAngleAddr

	return p
}

func (s *ObsAngleAddrContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsAngleAddrContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressParserLess, 0)
}

func (s *ObsAngleAddrContext) ObsRoute() IObsRouteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsRouteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsRouteContext)
}

func (s *ObsAngleAddrContext) AddrSpec() IAddrSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddrSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddrSpecContext)
}

func (s *ObsAngleAddrContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressParserGreater, 0)
}

func (s *ObsAngleAddrContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *ObsAngleAddrContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *ObsAngleAddrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsAngleAddrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsAngleAddrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsAngleAddr(s)
	}
}

func (s *ObsAngleAddrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsAngleAddr(s)
	}
}

func (p *AddressParser) ObsAngleAddr() (localctx IObsAngleAddrContext) {
	localctx = NewObsAngleAddrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AddressParserRULE_obsAngleAddr)
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
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(365)
			p.Cfws()
		}

	}
	{
		p.SetState(368)
		p.Match(AddressParserLess)
	}
	{
		p.SetState(369)
		p.ObsRoute()
	}
	{
		p.SetState(370)
		p.AddrSpec()
	}
	{
		p.SetState(371)
		p.Match(AddressParserGreater)
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(372)
			p.Cfws()
		}

	}

	return localctx
}

// IObsRouteContext is an interface to support dynamic dispatch.
type IObsRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsRouteContext differentiates from other interfaces.
	IsObsRouteContext()
}

type ObsRouteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsRouteContext() *ObsRouteContext {
	var p = new(ObsRouteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsRoute
	return p
}

func (*ObsRouteContext) IsObsRouteContext() {}

func NewObsRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsRouteContext {
	var p = new(ObsRouteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsRoute

	return p
}

func (s *ObsRouteContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsRouteContext) ObsDomainList() IObsDomainListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsDomainListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsDomainListContext)
}

func (s *ObsRouteContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressParserColon, 0)
}

func (s *ObsRouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsRouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsRouteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsRoute(s)
	}
}

func (s *ObsRouteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsRoute(s)
	}
}

func (p *AddressParser) ObsRoute() (localctx IObsRouteContext) {
	localctx = NewObsRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AddressParserRULE_obsRoute)

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
		p.SetState(375)
		p.ObsDomainList()
	}
	{
		p.SetState(376)
		p.Match(AddressParserColon)
	}

	return localctx
}

// IObsDomainListContext is an interface to support dynamic dispatch.
type IObsDomainListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsDomainListContext differentiates from other interfaces.
	IsObsDomainListContext()
}

type ObsDomainListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsDomainListContext() *ObsDomainListContext {
	var p = new(ObsDomainListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsDomainList
	return p
}

func (*ObsDomainListContext) IsObsDomainListContext() {}

func NewObsDomainListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsDomainListContext {
	var p = new(ObsDomainListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsDomainList

	return p
}

func (s *ObsDomainListContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsDomainListContext) AllAt() []antlr.TerminalNode {
	return s.GetTokens(AddressParserAt)
}

func (s *ObsDomainListContext) At(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserAt, i)
}

func (s *ObsDomainListContext) AllDomain() []IDomainContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDomainContext)(nil)).Elem())
	var tst = make([]IDomainContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDomainContext)
		}
	}

	return tst
}

func (s *ObsDomainListContext) Domain(i int) IDomainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDomainContext)
}

func (s *ObsDomainListContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *ObsDomainListContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *ObsDomainListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AddressParserComma)
}

func (s *ObsDomainListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserComma, i)
}

func (s *ObsDomainListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsDomainListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsDomainListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsDomainList(s)
	}
}

func (s *ObsDomainListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsDomainList(s)
	}
}

func (p *AddressParser) ObsDomainList() (localctx IObsDomainListContext) {
	localctx = NewObsDomainListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, AddressParserRULE_obsDomainList)
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
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens)|(1<<AddressParserComma))) != 0 {
		p.SetState(380)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AddressParserTAB, AddressParserCR, AddressParserSP, AddressParserLParens:
			{
				p.SetState(378)
				p.Cfws()
			}

		case AddressParserComma:
			{
				p.SetState(379)
				p.Match(AddressParserComma)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(385)
		p.Match(AddressParserAt)
	}
	{
		p.SetState(386)
		p.Domain()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(387)
			p.Match(AddressParserComma)
		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(388)
				p.Cfws()
			}

		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AddressParserAt {
			{
				p.SetState(391)
				p.Match(AddressParserAt)
			}
			{
				p.SetState(392)
				p.Domain()
			}

		}

		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObsMboxListContext is an interface to support dynamic dispatch.
type IObsMboxListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsMboxListContext differentiates from other interfaces.
	IsObsMboxListContext()
}

type ObsMboxListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsMboxListContext() *ObsMboxListContext {
	var p = new(ObsMboxListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsMboxList
	return p
}

func (*ObsMboxListContext) IsObsMboxListContext() {}

func NewObsMboxListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsMboxListContext {
	var p = new(ObsMboxListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsMboxList

	return p
}

func (s *ObsMboxListContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsMboxListContext) AllMailbox() []IMailboxContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMailboxContext)(nil)).Elem())
	var tst = make([]IMailboxContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMailboxContext)
		}
	}

	return tst
}

func (s *ObsMboxListContext) Mailbox(i int) IMailboxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMailboxContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMailboxContext)
}

func (s *ObsMboxListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AddressParserComma)
}

func (s *ObsMboxListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserComma, i)
}

func (s *ObsMboxListContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *ObsMboxListContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *ObsMboxListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsMboxListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsMboxListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsMboxList(s)
	}
}

func (s *ObsMboxListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsMboxList(s)
	}
}

func (p *AddressParser) ObsMboxList() (localctx IObsMboxListContext) {
	localctx = NewObsMboxListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, AddressParserRULE_obsMboxList)
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
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(401)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(400)
					p.Cfws()
				}

			}
			{
				p.SetState(403)
				p.Match(AddressParserComma)
			}

		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())
	}
	{
		p.SetState(409)
		p.Mailbox()
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(410)
			p.Match(AddressParserComma)
		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(411)
				p.Mailbox()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(412)
				p.Cfws()
			}

		}

		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObsAddrListContext is an interface to support dynamic dispatch.
type IObsAddrListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsAddrListContext differentiates from other interfaces.
	IsObsAddrListContext()
}

type ObsAddrListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsAddrListContext() *ObsAddrListContext {
	var p = new(ObsAddrListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsAddrList
	return p
}

func (*ObsAddrListContext) IsObsAddrListContext() {}

func NewObsAddrListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsAddrListContext {
	var p = new(ObsAddrListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsAddrList

	return p
}

func (s *ObsAddrListContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsAddrListContext) AllAddress() []IAddressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddressContext)(nil)).Elem())
	var tst = make([]IAddressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddressContext)
		}
	}

	return tst
}

func (s *ObsAddrListContext) Address(i int) IAddressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddressContext)
}

func (s *ObsAddrListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AddressParserComma)
}

func (s *ObsAddrListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserComma, i)
}

func (s *ObsAddrListContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *ObsAddrListContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *ObsAddrListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsAddrListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsAddrListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsAddrList(s)
	}
}

func (s *ObsAddrListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsAddrList(s)
	}
}

func (p *AddressParser) ObsAddrList() (localctx IObsAddrListContext) {
	localctx = NewObsAddrListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, AddressParserRULE_obsAddrList)
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
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(420)
					p.Cfws()
				}

			}
			{
				p.SetState(423)
				p.Match(AddressParserComma)
			}

		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext())
	}
	{
		p.SetState(429)
		p.Address()
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(430)
			p.Match(AddressParserComma)
		}
		p.SetState(433)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(431)
				p.Address()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(432)
				p.Cfws()
			}

		}

		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObsGroupListContext is an interface to support dynamic dispatch.
type IObsGroupListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsGroupListContext differentiates from other interfaces.
	IsObsGroupListContext()
}

type ObsGroupListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsGroupListContext() *ObsGroupListContext {
	var p = new(ObsGroupListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsGroupList
	return p
}

func (*ObsGroupListContext) IsObsGroupListContext() {}

func NewObsGroupListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsGroupListContext {
	var p = new(ObsGroupListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsGroupList

	return p
}

func (s *ObsGroupListContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsGroupListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AddressParserComma)
}

func (s *ObsGroupListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserComma, i)
}

func (s *ObsGroupListContext) AllCfws() []ICfwsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICfwsContext)(nil)).Elem())
	var tst = make([]ICfwsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICfwsContext)
		}
	}

	return tst
}

func (s *ObsGroupListContext) Cfws(i int) ICfwsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICfwsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICfwsContext)
}

func (s *ObsGroupListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsGroupListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsGroupListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsGroupList(s)
	}
}

func (s *ObsGroupListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsGroupList(s)
	}
}

func (p *AddressParser) ObsGroupList() (localctx IObsGroupListContext) {
	localctx = NewObsGroupListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, AddressParserRULE_obsGroupList)
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
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(441)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(440)
					p.Cfws()
				}

			}
			{
				p.SetState(443)
				p.Match(AddressParserComma)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext())
	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(448)
			p.Cfws()
		}

	}

	return localctx
}

// IEncodedWordContext is an interface to support dynamic dispatch.
type IEncodedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncodedWordContext differentiates from other interfaces.
	IsEncodedWordContext()
}

type EncodedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncodedWordContext() *EncodedWordContext {
	var p = new(EncodedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_encodedWord
	return p
}

func (*EncodedWordContext) IsEncodedWordContext() {}

func NewEncodedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncodedWordContext {
	var p = new(EncodedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_encodedWord

	return p
}

func (s *EncodedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *EncodedWordContext) AllEqual() []antlr.TerminalNode {
	return s.GetTokens(AddressParserEqual)
}

func (s *EncodedWordContext) Equal(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserEqual, i)
}

func (s *EncodedWordContext) AllQuestion() []antlr.TerminalNode {
	return s.GetTokens(AddressParserQuestion)
}

func (s *EncodedWordContext) Question(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserQuestion, i)
}

func (s *EncodedWordContext) Charset() ICharsetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharsetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharsetContext)
}

func (s *EncodedWordContext) Encoding() IEncodingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncodingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEncodingContext)
}

func (s *EncodedWordContext) EncodedText() IEncodedTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncodedTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEncodedTextContext)
}

func (s *EncodedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncodedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncodedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterEncodedWord(s)
	}
}

func (s *EncodedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitEncodedWord(s)
	}
}

func (p *AddressParser) EncodedWord() (localctx IEncodedWordContext) {
	localctx = NewEncodedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, AddressParserRULE_encodedWord)

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
		p.SetState(451)
		p.Match(AddressParserEqual)
	}
	{
		p.SetState(452)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(453)
		p.Charset()
	}
	{
		p.SetState(454)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(455)
		p.Encoding()
	}
	{
		p.SetState(456)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(457)
		p.EncodedText()
	}
	{
		p.SetState(458)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(459)
		p.Match(AddressParserEqual)
	}

	return localctx
}

// ICharsetContext is an interface to support dynamic dispatch.
type ICharsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharsetContext differentiates from other interfaces.
	IsCharsetContext()
}

type CharsetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharsetContext() *CharsetContext {
	var p = new(CharsetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_charset
	return p
}

func (*CharsetContext) IsCharsetContext() {}

func NewCharsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharsetContext {
	var p = new(CharsetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_charset

	return p
}

func (s *CharsetContext) GetParser() antlr.Parser { return s.parser }

func (s *CharsetContext) Token() ITokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITokenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITokenContext)
}

func (s *CharsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterCharset(s)
	}
}

func (s *CharsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitCharset(s)
	}
}

func (p *AddressParser) Charset() (localctx ICharsetContext) {
	localctx = NewCharsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, AddressParserRULE_charset)

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
		p.SetState(461)
		p.Token()
	}

	return localctx
}

// IEncodingContext is an interface to support dynamic dispatch.
type IEncodingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncodingContext differentiates from other interfaces.
	IsEncodingContext()
}

type EncodingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncodingContext() *EncodingContext {
	var p = new(EncodingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_encoding
	return p
}

func (*EncodingContext) IsEncodingContext() {}

func NewEncodingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncodingContext {
	var p = new(EncodingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_encoding

	return p
}

func (s *EncodingContext) GetParser() antlr.Parser { return s.parser }

func (s *EncodingContext) Token() ITokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITokenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITokenContext)
}

func (s *EncodingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncodingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncodingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterEncoding(s)
	}
}

func (s *EncodingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitEncoding(s)
	}
}

func (p *AddressParser) Encoding() (localctx IEncodingContext) {
	localctx = NewEncodingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, AddressParserRULE_encoding)

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
		p.SetState(463)
		p.Token()
	}

	return localctx
}

// ITokenContext is an interface to support dynamic dispatch.
type ITokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTokenContext differentiates from other interfaces.
	IsTokenContext()
}

type TokenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokenContext() *TokenContext {
	var p = new(TokenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_token
	return p
}

func (*TokenContext) IsTokenContext() {}

func NewTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenContext {
	var p = new(TokenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_token

	return p
}

func (s *TokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenContext) AllTokenChar() []ITokenCharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITokenCharContext)(nil)).Elem())
	var tst = make([]ITokenCharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITokenCharContext)
		}
	}

	return tst
}

func (s *TokenContext) TokenChar(i int) ITokenCharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITokenCharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITokenCharContext)
}

func (s *TokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterToken(s)
	}
}

func (s *TokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitToken(s)
	}
}

func (p *AddressParser) Token() (localctx ITokenContext) {
	localctx = NewTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, AddressParserRULE_token)
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
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserDigit)|(1<<AddressParserAlphaUpper)|(1<<AddressParserBackslash))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0) {
		{
			p.SetState(465)
			p.TokenChar()
		}

		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITokenCharContext is an interface to support dynamic dispatch.
type ITokenCharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTokenCharContext differentiates from other interfaces.
	IsTokenCharContext()
}

type TokenCharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokenCharContext() *TokenCharContext {
	var p = new(TokenCharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_tokenChar
	return p
}

func (*TokenCharContext) IsTokenCharContext() {}

func NewTokenCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenCharContext {
	var p = new(TokenCharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_tokenChar

	return p
}

func (s *TokenCharContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenCharContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressParserExclamation, 0)
}

func (s *TokenCharContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressParserHash, 0)
}

func (s *TokenCharContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressParserDollar, 0)
}

func (s *TokenCharContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressParserPercent, 0)
}

func (s *TokenCharContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressParserAmpersand, 0)
}

func (s *TokenCharContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserSQuote, 0)
}

func (s *TokenCharContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressParserAsterisk, 0)
}

func (s *TokenCharContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressParserPlus, 0)
}

func (s *TokenCharContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressParserMinus, 0)
}

func (s *TokenCharContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressParserDigit, 0)
}

func (s *TokenCharContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaUpper, 0)
}

func (s *TokenCharContext) Backslash() antlr.TerminalNode {
	return s.GetToken(AddressParserBackslash, 0)
}

func (s *TokenCharContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressParserCaret, 0)
}

func (s *TokenCharContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressParserUnderscore, 0)
}

func (s *TokenCharContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressParserBacktick, 0)
}

func (s *TokenCharContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaLower, 0)
}

func (s *TokenCharContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserLCurly, 0)
}

func (s *TokenCharContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressParserPipe, 0)
}

func (s *TokenCharContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserRCurly, 0)
}

func (s *TokenCharContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressParserTilde, 0)
}

func (s *TokenCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokenCharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokenCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterTokenChar(s)
	}
}

func (s *TokenCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitTokenChar(s)
	}
}

func (p *AddressParser) TokenChar() (localctx ITokenCharContext) {
	localctx = NewTokenCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, AddressParserRULE_tokenChar)
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
		p.SetState(470)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserDigit)|(1<<AddressParserAlphaUpper)|(1<<AddressParserBackslash))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEncodedTextContext is an interface to support dynamic dispatch.
type IEncodedTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncodedTextContext differentiates from other interfaces.
	IsEncodedTextContext()
}

type EncodedTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncodedTextContext() *EncodedTextContext {
	var p = new(EncodedTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_encodedText
	return p
}

func (*EncodedTextContext) IsEncodedTextContext() {}

func NewEncodedTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncodedTextContext {
	var p = new(EncodedTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_encodedText

	return p
}

func (s *EncodedTextContext) GetParser() antlr.Parser { return s.parser }

func (s *EncodedTextContext) AllEncodedChar() []IEncodedCharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEncodedCharContext)(nil)).Elem())
	var tst = make([]IEncodedCharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEncodedCharContext)
		}
	}

	return tst
}

func (s *EncodedTextContext) EncodedChar(i int) IEncodedCharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncodedCharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEncodedCharContext)
}

func (s *EncodedTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncodedTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncodedTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterEncodedText(s)
	}
}

func (s *EncodedTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitEncodedText(s)
	}
}

func (p *AddressParser) EncodedText() (localctx IEncodedTextContext) {
	localctx = NewEncodedTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, AddressParserRULE_encodedText)
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
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater)|(1<<AddressParserAt)|(1<<AddressParserAlphaUpper)|(1<<AddressParserLBracket)|(1<<AddressParserBackslash)|(1<<AddressParserRBracket))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0) {
		{
			p.SetState(472)
			p.EncodedChar()
		}

		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEncodedCharContext is an interface to support dynamic dispatch.
type IEncodedCharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncodedCharContext differentiates from other interfaces.
	IsEncodedCharContext()
}

type EncodedCharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncodedCharContext() *EncodedCharContext {
	var p = new(EncodedCharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_encodedChar
	return p
}

func (*EncodedCharContext) IsEncodedCharContext() {}

func NewEncodedCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncodedCharContext {
	var p = new(EncodedCharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_encodedChar

	return p
}

func (s *EncodedCharContext) GetParser() antlr.Parser { return s.parser }

func (s *EncodedCharContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(AddressParserExclamation, 0)
}

func (s *EncodedCharContext) DQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserDQuote, 0)
}

func (s *EncodedCharContext) Hash() antlr.TerminalNode {
	return s.GetToken(AddressParserHash, 0)
}

func (s *EncodedCharContext) Dollar() antlr.TerminalNode {
	return s.GetToken(AddressParserDollar, 0)
}

func (s *EncodedCharContext) Percent() antlr.TerminalNode {
	return s.GetToken(AddressParserPercent, 0)
}

func (s *EncodedCharContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(AddressParserAmpersand, 0)
}

func (s *EncodedCharContext) SQuote() antlr.TerminalNode {
	return s.GetToken(AddressParserSQuote, 0)
}

func (s *EncodedCharContext) LParens() antlr.TerminalNode {
	return s.GetToken(AddressParserLParens, 0)
}

func (s *EncodedCharContext) RParens() antlr.TerminalNode {
	return s.GetToken(AddressParserRParens, 0)
}

func (s *EncodedCharContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(AddressParserAsterisk, 0)
}

func (s *EncodedCharContext) Plus() antlr.TerminalNode {
	return s.GetToken(AddressParserPlus, 0)
}

func (s *EncodedCharContext) Comma() antlr.TerminalNode {
	return s.GetToken(AddressParserComma, 0)
}

func (s *EncodedCharContext) Minus() antlr.TerminalNode {
	return s.GetToken(AddressParserMinus, 0)
}

func (s *EncodedCharContext) Period() antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, 0)
}

func (s *EncodedCharContext) Slash() antlr.TerminalNode {
	return s.GetToken(AddressParserSlash, 0)
}

func (s *EncodedCharContext) Digit() antlr.TerminalNode {
	return s.GetToken(AddressParserDigit, 0)
}

func (s *EncodedCharContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressParserColon, 0)
}

func (s *EncodedCharContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AddressParserSemicolon, 0)
}

func (s *EncodedCharContext) Less() antlr.TerminalNode {
	return s.GetToken(AddressParserLess, 0)
}

func (s *EncodedCharContext) Equal() antlr.TerminalNode {
	return s.GetToken(AddressParserEqual, 0)
}

func (s *EncodedCharContext) Greater() antlr.TerminalNode {
	return s.GetToken(AddressParserGreater, 0)
}

func (s *EncodedCharContext) At() antlr.TerminalNode {
	return s.GetToken(AddressParserAt, 0)
}

func (s *EncodedCharContext) AlphaUpper() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaUpper, 0)
}

func (s *EncodedCharContext) LBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserLBracket, 0)
}

func (s *EncodedCharContext) Backslash() antlr.TerminalNode {
	return s.GetToken(AddressParserBackslash, 0)
}

func (s *EncodedCharContext) RBracket() antlr.TerminalNode {
	return s.GetToken(AddressParserRBracket, 0)
}

func (s *EncodedCharContext) Caret() antlr.TerminalNode {
	return s.GetToken(AddressParserCaret, 0)
}

func (s *EncodedCharContext) Underscore() antlr.TerminalNode {
	return s.GetToken(AddressParserUnderscore, 0)
}

func (s *EncodedCharContext) Backtick() antlr.TerminalNode {
	return s.GetToken(AddressParserBacktick, 0)
}

func (s *EncodedCharContext) AlphaLower() antlr.TerminalNode {
	return s.GetToken(AddressParserAlphaLower, 0)
}

func (s *EncodedCharContext) LCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserLCurly, 0)
}

func (s *EncodedCharContext) Pipe() antlr.TerminalNode {
	return s.GetToken(AddressParserPipe, 0)
}

func (s *EncodedCharContext) RCurly() antlr.TerminalNode {
	return s.GetToken(AddressParserRCurly, 0)
}

func (s *EncodedCharContext) Tilde() antlr.TerminalNode {
	return s.GetToken(AddressParserTilde, 0)
}

func (s *EncodedCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncodedCharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncodedCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterEncodedChar(s)
	}
}

func (s *EncodedCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitEncodedChar(s)
	}
}

func (p *AddressParser) EncodedChar() (localctx IEncodedCharContext) {
	localctx = NewEncodedCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, AddressParserRULE_encodedChar)
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
		p.SetState(477)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater)|(1<<AddressParserAt)|(1<<AddressParserAlphaUpper)|(1<<AddressParserLBracket)|(1<<AddressParserBackslash)|(1<<AddressParserRBracket))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0)) {
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
	p.EnterRule(localctx, 88, AddressParserRULE_crlf)

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
		p.SetState(479)
		p.Match(AddressParserCR)
	}
	{
		p.SetState(480)
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
	p.EnterRule(localctx, 90, AddressParserRULE_wsp)
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
		p.SetState(482)
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
	p.EnterRule(localctx, 92, AddressParserRULE_vchar)
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
		p.SetState(484)
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
