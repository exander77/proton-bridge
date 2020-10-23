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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 604,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 3, 2,
	3, 2, 5, 2, 111, 10, 2, 3, 3, 3, 3, 3, 3, 5, 3, 116, 10, 3, 3, 4, 7, 4,
	119, 10, 4, 12, 4, 14, 4, 122, 11, 4, 3, 4, 5, 4, 125, 10, 4, 3, 4, 6,
	4, 128, 10, 4, 13, 4, 14, 4, 129, 3, 4, 5, 4, 133, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 169, 10, 5, 3, 6, 3,
	6, 3, 6, 5, 6, 174, 10, 6, 3, 7, 3, 7, 5, 7, 178, 10, 7, 3, 7, 7, 7, 181,
	10, 7, 12, 7, 14, 7, 184, 11, 7, 3, 7, 5, 7, 187, 10, 7, 3, 7, 3, 7, 3,
	8, 5, 8, 192, 10, 8, 3, 8, 6, 8, 195, 10, 8, 13, 8, 14, 8, 196, 3, 8, 5,
	8, 200, 10, 8, 3, 8, 5, 8, 203, 10, 8, 3, 9, 3, 9, 3, 10, 6, 10, 208, 10,
	10, 13, 10, 14, 10, 209, 3, 11, 6, 11, 213, 10, 11, 13, 11, 14, 11, 214,
	3, 11, 3, 11, 6, 11, 219, 10, 11, 13, 11, 14, 11, 220, 7, 11, 223, 10,
	11, 12, 11, 14, 11, 226, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 263,
	10, 12, 3, 13, 3, 13, 5, 13, 267, 10, 13, 3, 14, 5, 14, 270, 10, 14, 3,
	14, 7, 14, 273, 10, 14, 12, 14, 14, 14, 276, 11, 14, 3, 15, 3, 15, 3, 15,
	5, 15, 281, 10, 15, 3, 15, 3, 15, 3, 16, 5, 16, 286, 10, 16, 3, 16, 3,
	16, 5, 16, 290, 10, 16, 3, 16, 5, 16, 293, 10, 16, 3, 16, 3, 16, 5, 16,
	297, 10, 16, 3, 16, 5, 16, 300, 10, 16, 3, 16, 3, 16, 5, 16, 304, 10, 16,
	5, 16, 306, 10, 16, 3, 17, 3, 17, 5, 17, 310, 10, 17, 3, 18, 3, 18, 5,
	18, 314, 10, 18, 3, 19, 5, 19, 317, 10, 19, 3, 19, 3, 19, 3, 20, 5, 20,
	322, 10, 20, 3, 20, 3, 20, 5, 20, 326, 10, 20, 3, 20, 3, 20, 5, 20, 330,
	10, 20, 3, 20, 5, 20, 333, 10, 20, 3, 21, 3, 21, 3, 21, 5, 21, 338, 10,
	21, 3, 21, 3, 21, 5, 21, 342, 10, 21, 3, 22, 6, 22, 345, 10, 22, 13, 22,
	14, 22, 346, 3, 22, 5, 22, 350, 10, 22, 3, 23, 3, 23, 3, 23, 7, 23, 355,
	10, 23, 12, 23, 14, 23, 358, 11, 23, 3, 23, 5, 23, 361, 10, 23, 3, 24,
	3, 24, 3, 24, 7, 24, 366, 10, 24, 12, 24, 14, 24, 369, 11, 24, 3, 24, 5,
	24, 372, 10, 24, 3, 25, 3, 25, 3, 25, 5, 25, 377, 10, 25, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 27, 5, 27, 384, 10, 27, 3, 27, 3, 27, 5, 27, 388, 10,
	27, 3, 27, 5, 27, 391, 10, 27, 3, 27, 3, 27, 5, 27, 395, 10, 27, 3, 27,
	5, 27, 398, 10, 27, 3, 28, 5, 28, 401, 10, 28, 3, 28, 3, 28, 5, 28, 405,
	10, 28, 3, 28, 5, 28, 408, 10, 28, 3, 28, 3, 28, 5, 28, 412, 10, 28, 3,
	28, 5, 28, 415, 10, 28, 3, 29, 3, 29, 5, 29, 419, 10, 29, 3, 29, 7, 29,
	422, 10, 29, 12, 29, 14, 29, 425, 11, 29, 3, 29, 5, 29, 428, 10, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32,
	440, 10, 32, 12, 32, 14, 32, 443, 11, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 454, 10, 35, 3, 36, 6, 36, 457,
	10, 36, 13, 36, 14, 36, 458, 3, 36, 3, 36, 6, 36, 463, 10, 36, 13, 36,
	14, 36, 464, 3, 37, 5, 37, 468, 10, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 5, 37, 475, 10, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 7, 39, 482,
	10, 39, 12, 39, 14, 39, 485, 11, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39,
	491, 10, 39, 3, 39, 3, 39, 5, 39, 495, 10, 39, 7, 39, 497, 10, 39, 12,
	39, 14, 39, 500, 11, 39, 3, 40, 5, 40, 503, 10, 40, 3, 40, 7, 40, 506,
	10, 40, 12, 40, 14, 40, 509, 11, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40,
	515, 10, 40, 7, 40, 517, 10, 40, 12, 40, 14, 40, 520, 11, 40, 3, 41, 5,
	41, 523, 10, 41, 3, 41, 7, 41, 526, 10, 41, 12, 41, 14, 41, 529, 11, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 535, 10, 41, 7, 41, 537, 10, 41, 12,
	41, 14, 41, 540, 11, 41, 3, 42, 5, 42, 543, 10, 42, 3, 42, 6, 42, 546,
	10, 42, 13, 42, 14, 42, 547, 3, 42, 5, 42, 551, 10, 42, 3, 43, 3, 43, 3,
	43, 7, 43, 556, 10, 43, 12, 43, 14, 43, 559, 11, 43, 3, 44, 3, 44, 3, 44,
	7, 44, 564, 10, 44, 12, 44, 14, 44, 567, 11, 44, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47,
	3, 48, 6, 48, 584, 10, 48, 13, 48, 14, 48, 585, 3, 49, 3, 49, 3, 50, 6,
	50, 591, 10, 50, 13, 50, 14, 50, 592, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52,
	3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 2, 2, 55, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
	54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
	90, 92, 94, 96, 98, 100, 102, 104, 106, 2, 9, 12, 2, 12, 12, 14, 18, 21,
	22, 24, 24, 26, 27, 31, 31, 33, 33, 35, 35, 39, 46, 48, 48, 5, 2, 12, 35,
	39, 46, 48, 48, 6, 2, 4, 4, 7, 8, 10, 10, 47, 47, 10, 2, 12, 12, 14, 18,
	21, 22, 24, 24, 27, 27, 35, 35, 37, 37, 39, 46, 4, 2, 12, 32, 34, 46, 4,
	2, 5, 5, 11, 11, 4, 2, 12, 46, 48, 48, 2, 712, 2, 110, 3, 2, 2, 2, 4, 115,
	3, 2, 2, 2, 6, 132, 3, 2, 2, 2, 8, 168, 3, 2, 2, 2, 10, 173, 3, 2, 2, 2,
	12, 175, 3, 2, 2, 2, 14, 202, 3, 2, 2, 2, 16, 204, 3, 2, 2, 2, 18, 207,
	3, 2, 2, 2, 20, 212, 3, 2, 2, 2, 22, 262, 3, 2, 2, 2, 24, 266, 3, 2, 2,
	2, 26, 274, 3, 2, 2, 2, 28, 277, 3, 2, 2, 2, 30, 305, 3, 2, 2, 2, 32, 309,
	3, 2, 2, 2, 34, 313, 3, 2, 2, 2, 36, 316, 3, 2, 2, 2, 38, 332, 3, 2, 2,
	2, 40, 334, 3, 2, 2, 2, 42, 349, 3, 2, 2, 2, 44, 360, 3, 2, 2, 2, 46, 371,
	3, 2, 2, 2, 48, 376, 3, 2, 2, 2, 50, 378, 3, 2, 2, 2, 52, 397, 3, 2, 2,
	2, 54, 414, 3, 2, 2, 2, 56, 416, 3, 2, 2, 2, 58, 431, 3, 2, 2, 2, 60, 433,
	3, 2, 2, 2, 62, 435, 3, 2, 2, 2, 64, 444, 3, 2, 2, 2, 66, 446, 3, 2, 2,
	2, 68, 448, 3, 2, 2, 2, 70, 456, 3, 2, 2, 2, 72, 467, 3, 2, 2, 2, 74, 476,
	3, 2, 2, 2, 76, 483, 3, 2, 2, 2, 78, 507, 3, 2, 2, 2, 80, 527, 3, 2, 2,
	2, 82, 545, 3, 2, 2, 2, 84, 552, 3, 2, 2, 2, 86, 560, 3, 2, 2, 2, 88, 568,
	3, 2, 2, 2, 90, 578, 3, 2, 2, 2, 92, 580, 3, 2, 2, 2, 94, 583, 3, 2, 2,
	2, 96, 587, 3, 2, 2, 2, 98, 590, 3, 2, 2, 2, 100, 594, 3, 2, 2, 2, 102,
	596, 3, 2, 2, 2, 104, 599, 3, 2, 2, 2, 106, 601, 3, 2, 2, 2, 108, 111,
	5, 106, 54, 2, 109, 111, 5, 104, 53, 2, 110, 108, 3, 2, 2, 2, 110, 109,
	3, 2, 2, 2, 111, 3, 3, 2, 2, 2, 112, 113, 7, 37, 2, 2, 113, 116, 5, 2,
	2, 2, 114, 116, 5, 68, 35, 2, 115, 112, 3, 2, 2, 2, 115, 114, 3, 2, 2,
	2, 116, 5, 3, 2, 2, 2, 117, 119, 5, 104, 53, 2, 118, 117, 3, 2, 2, 2, 119,
	122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123,
	3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 125, 5, 102, 52, 2, 124, 120, 3,
	2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2, 126, 128, 5, 104,
	53, 2, 127, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2,
	129, 130, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 133, 5, 70, 36, 2, 132,
	124, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 133, 7, 3, 2, 2, 2, 134, 169, 7,
	12, 2, 2, 135, 169, 7, 13, 2, 2, 136, 169, 7, 14, 2, 2, 137, 169, 7, 15,
	2, 2, 138, 169, 7, 16, 2, 2, 139, 169, 7, 17, 2, 2, 140, 169, 7, 18, 2,
	2, 141, 169, 7, 21, 2, 2, 142, 169, 7, 22, 2, 2, 143, 169, 7, 23, 2, 2,
	144, 169, 7, 24, 2, 2, 145, 169, 7, 25, 2, 2, 146, 169, 7, 26, 2, 2, 147,
	169, 7, 27, 2, 2, 148, 169, 7, 28, 2, 2, 149, 169, 7, 29, 2, 2, 150, 169,
	7, 30, 2, 2, 151, 169, 7, 31, 2, 2, 152, 169, 7, 32, 2, 2, 153, 169, 7,
	33, 2, 2, 154, 169, 7, 34, 2, 2, 155, 169, 7, 35, 2, 2, 156, 169, 7, 36,
	2, 2, 157, 169, 7, 38, 2, 2, 158, 169, 7, 39, 2, 2, 159, 169, 7, 40, 2,
	2, 160, 169, 7, 41, 2, 2, 161, 169, 7, 42, 2, 2, 162, 169, 7, 43, 2, 2,
	163, 169, 7, 44, 2, 2, 164, 169, 7, 45, 2, 2, 165, 169, 7, 46, 2, 2, 166,
	169, 5, 64, 33, 2, 167, 169, 7, 48, 2, 2, 168, 134, 3, 2, 2, 2, 168, 135,
	3, 2, 2, 2, 168, 136, 3, 2, 2, 2, 168, 137, 3, 2, 2, 2, 168, 138, 3, 2,
	2, 2, 168, 139, 3, 2, 2, 2, 168, 140, 3, 2, 2, 2, 168, 141, 3, 2, 2, 2,
	168, 142, 3, 2, 2, 2, 168, 143, 3, 2, 2, 2, 168, 144, 3, 2, 2, 2, 168,
	145, 3, 2, 2, 2, 168, 146, 3, 2, 2, 2, 168, 147, 3, 2, 2, 2, 168, 148,
	3, 2, 2, 2, 168, 149, 3, 2, 2, 2, 168, 150, 3, 2, 2, 2, 168, 151, 3, 2,
	2, 2, 168, 152, 3, 2, 2, 2, 168, 153, 3, 2, 2, 2, 168, 154, 3, 2, 2, 2,
	168, 155, 3, 2, 2, 2, 168, 156, 3, 2, 2, 2, 168, 157, 3, 2, 2, 2, 168,
	158, 3, 2, 2, 2, 168, 159, 3, 2, 2, 2, 168, 160, 3, 2, 2, 2, 168, 161,
	3, 2, 2, 2, 168, 162, 3, 2, 2, 2, 168, 163, 3, 2, 2, 2, 168, 164, 3, 2,
	2, 2, 168, 165, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 167, 3, 2, 2, 2,
	169, 9, 3, 2, 2, 2, 170, 174, 5, 8, 5, 2, 171, 174, 5, 4, 3, 2, 172, 174,
	5, 12, 7, 2, 173, 170, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 172, 3, 2,
	2, 2, 174, 11, 3, 2, 2, 2, 175, 182, 7, 19, 2, 2, 176, 178, 5, 6, 4, 2,
	177, 176, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179,
	181, 5, 10, 6, 2, 180, 177, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180,
	3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184, 182, 3, 2,
	2, 2, 185, 187, 5, 6, 4, 2, 186, 185, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2,
	187, 188, 3, 2, 2, 2, 188, 189, 7, 20, 2, 2, 189, 13, 3, 2, 2, 2, 190,
	192, 5, 6, 4, 2, 191, 190, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 193,
	3, 2, 2, 2, 193, 195, 5, 12, 7, 2, 194, 191, 3, 2, 2, 2, 195, 196, 3, 2,
	2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 199, 3, 2, 2, 2,
	198, 200, 5, 6, 4, 2, 199, 198, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200,
	203, 3, 2, 2, 2, 201, 203, 5, 6, 4, 2, 202, 194, 3, 2, 2, 2, 202, 201,
	3, 2, 2, 2, 203, 15, 3, 2, 2, 2, 204, 205, 9, 2, 2, 2, 205, 17, 3, 2, 2,
	2, 206, 208, 5, 16, 9, 2, 207, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209,
	207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 19, 3, 2, 2, 2, 211, 213, 5,
	16, 9, 2, 212, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 212, 3, 2, 2,
	2, 214, 215, 3, 2, 2, 2, 215, 224, 3, 2, 2, 2, 216, 218, 7, 25, 2, 2, 217,
	219, 5, 16, 9, 2, 218, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 218,
	3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 223, 3, 2, 2, 2, 222, 216, 3, 2,
	2, 2, 223, 226, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2,
	225, 21, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 227, 263, 7, 12, 2, 2, 228,
	263, 7, 14, 2, 2, 229, 263, 7, 15, 2, 2, 230, 263, 7, 16, 2, 2, 231, 263,
	7, 17, 2, 2, 232, 263, 7, 18, 2, 2, 233, 263, 7, 19, 2, 2, 234, 263, 7,
	20, 2, 2, 235, 263, 7, 21, 2, 2, 236, 263, 7, 22, 2, 2, 237, 263, 7, 23,
	2, 2, 238, 263, 7, 24, 2, 2, 239, 263, 7, 25, 2, 2, 240, 263, 7, 26, 2,
	2, 241, 263, 7, 27, 2, 2, 242, 263, 7, 28, 2, 2, 243, 263, 7, 29, 2, 2,
	244, 263, 7, 30, 2, 2, 245, 263, 7, 31, 2, 2, 246, 263, 7, 32, 2, 2, 247,
	263, 7, 33, 2, 2, 248, 263, 7, 34, 2, 2, 249, 263, 7, 35, 2, 2, 250, 263,
	7, 36, 2, 2, 251, 263, 7, 38, 2, 2, 252, 263, 7, 39, 2, 2, 253, 263, 7,
	40, 2, 2, 254, 263, 7, 41, 2, 2, 255, 263, 7, 42, 2, 2, 256, 263, 7, 43,
	2, 2, 257, 263, 7, 44, 2, 2, 258, 263, 7, 45, 2, 2, 259, 263, 7, 46, 2,
	2, 260, 263, 5, 66, 34, 2, 261, 263, 7, 48, 2, 2, 262, 227, 3, 2, 2, 2,
	262, 228, 3, 2, 2, 2, 262, 229, 3, 2, 2, 2, 262, 230, 3, 2, 2, 2, 262,
	231, 3, 2, 2, 2, 262, 232, 3, 2, 2, 2, 262, 233, 3, 2, 2, 2, 262, 234,
	3, 2, 2, 2, 262, 235, 3, 2, 2, 2, 262, 236, 3, 2, 2, 2, 262, 237, 3, 2,
	2, 2, 262, 238, 3, 2, 2, 2, 262, 239, 3, 2, 2, 2, 262, 240, 3, 2, 2, 2,
	262, 241, 3, 2, 2, 2, 262, 242, 3, 2, 2, 2, 262, 243, 3, 2, 2, 2, 262,
	244, 3, 2, 2, 2, 262, 245, 3, 2, 2, 2, 262, 246, 3, 2, 2, 2, 262, 247,
	3, 2, 2, 2, 262, 248, 3, 2, 2, 2, 262, 249, 3, 2, 2, 2, 262, 250, 3, 2,
	2, 2, 262, 251, 3, 2, 2, 2, 262, 252, 3, 2, 2, 2, 262, 253, 3, 2, 2, 2,
	262, 254, 3, 2, 2, 2, 262, 255, 3, 2, 2, 2, 262, 256, 3, 2, 2, 2, 262,
	257, 3, 2, 2, 2, 262, 258, 3, 2, 2, 2, 262, 259, 3, 2, 2, 2, 262, 260,
	3, 2, 2, 2, 262, 261, 3, 2, 2, 2, 263, 23, 3, 2, 2, 2, 264, 267, 5, 22,
	12, 2, 265, 267, 5, 4, 3, 2, 266, 264, 3, 2, 2, 2, 266, 265, 3, 2, 2, 2,
	267, 25, 3, 2, 2, 2, 268, 270, 5, 6, 4, 2, 269, 268, 3, 2, 2, 2, 269, 270,
	3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 273, 5, 24, 13, 2, 272, 269, 3,
	2, 2, 2, 273, 276, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2,
	2, 275, 27, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 277, 278, 7, 13, 2, 2, 278,
	280, 5, 26, 14, 2, 279, 281, 5, 6, 4, 2, 280, 279, 3, 2, 2, 2, 280, 281,
	3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 283, 7, 13, 2, 2, 283, 29, 3, 2,
	2, 2, 284, 286, 5, 14, 8, 2, 285, 284, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2,
	286, 287, 3, 2, 2, 2, 287, 289, 5, 88, 45, 2, 288, 290, 5, 14, 8, 2, 289,
	288, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 306, 3, 2, 2, 2, 291, 293,
	5, 14, 8, 2, 292, 291, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 294, 3, 2,
	2, 2, 294, 296, 5, 18, 10, 2, 295, 297, 5, 14, 8, 2, 296, 295, 3, 2, 2,
	2, 296, 297, 3, 2, 2, 2, 297, 306, 3, 2, 2, 2, 298, 300, 5, 14, 8, 2, 299,
	298, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 303,
	5, 28, 15, 2, 302, 304, 5, 14, 8, 2, 303, 302, 3, 2, 2, 2, 303, 304, 3,
	2, 2, 2, 304, 306, 3, 2, 2, 2, 305, 285, 3, 2, 2, 2, 305, 292, 3, 2, 2,
	2, 305, 299, 3, 2, 2, 2, 306, 31, 3, 2, 2, 2, 307, 310, 5, 34, 18, 2, 308,
	310, 5, 40, 21, 2, 309, 307, 3, 2, 2, 2, 309, 308, 3, 2, 2, 2, 310, 33,
	3, 2, 2, 2, 311, 314, 5, 36, 19, 2, 312, 314, 5, 50, 26, 2, 313, 311, 3,
	2, 2, 2, 313, 312, 3, 2, 2, 2, 314, 35, 3, 2, 2, 2, 315, 317, 5, 42, 22,
	2, 316, 315, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318,
	319, 5, 38, 20, 2, 319, 37, 3, 2, 2, 2, 320, 322, 5, 14, 8, 2, 321, 320,
	3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 325, 7, 30,
	2, 2, 324, 326, 5, 50, 26, 2, 325, 324, 3, 2, 2, 2, 325, 326, 3, 2, 2,
	2, 326, 327, 3, 2, 2, 2, 327, 329, 7, 32, 2, 2, 328, 330, 5, 14, 8, 2,
	329, 328, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 333, 3, 2, 2, 2, 331,
	333, 5, 72, 37, 2, 332, 321, 3, 2, 2, 2, 332, 331, 3, 2, 2, 2, 333, 39,
	3, 2, 2, 2, 334, 335, 5, 42, 22, 2, 335, 337, 7, 28, 2, 2, 336, 338, 5,
	48, 25, 2, 337, 336, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 339, 3, 2,
	2, 2, 339, 341, 7, 29, 2, 2, 340, 342, 5, 14, 8, 2, 341, 340, 3, 2, 2,
	2, 341, 342, 3, 2, 2, 2, 342, 41, 3, 2, 2, 2, 343, 345, 5, 30, 16, 2, 344,
	343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346, 347,
	3, 2, 2, 2, 347, 350, 3, 2, 2, 2, 348, 350, 5, 62, 32, 2, 349, 344, 3,
	2, 2, 2, 349, 348, 3, 2, 2, 2, 350, 43, 3, 2, 2, 2, 351, 356, 5, 34, 18,
	2, 352, 353, 7, 23, 2, 2, 353, 355, 5, 34, 18, 2, 354, 352, 3, 2, 2, 2,
	355, 358, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357,
	361, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 359, 361, 5, 78, 40, 2, 360, 351,
	3, 2, 2, 2, 360, 359, 3, 2, 2, 2, 361, 45, 3, 2, 2, 2, 362, 367, 5, 32,
	17, 2, 363, 364, 7, 23, 2, 2, 364, 366, 5, 32, 17, 2, 365, 363, 3, 2, 2,
	2, 366, 369, 3, 2, 2, 2, 367, 365, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368,
	372, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 370, 372, 5, 80, 41, 2, 371, 362,
	3, 2, 2, 2, 371, 370, 3, 2, 2, 2, 372, 47, 3, 2, 2, 2, 373, 377, 5, 44,
	23, 2, 374, 377, 5, 14, 8, 2, 375, 377, 5, 82, 42, 2, 376, 373, 3, 2, 2,
	2, 376, 374, 3, 2, 2, 2, 376, 375, 3, 2, 2, 2, 377, 49, 3, 2, 2, 2, 378,
	379, 5, 52, 27, 2, 379, 380, 7, 34, 2, 2, 380, 381, 5, 54, 28, 2, 381,
	51, 3, 2, 2, 2, 382, 384, 5, 14, 8, 2, 383, 382, 3, 2, 2, 2, 383, 384,
	3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 387, 5, 20, 11, 2, 386, 388, 5,
	14, 8, 2, 387, 386, 3, 2, 2, 2, 387, 388, 3, 2, 2, 2, 388, 398, 3, 2, 2,
	2, 389, 391, 5, 14, 8, 2, 390, 389, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391,
	392, 3, 2, 2, 2, 392, 394, 5, 28, 15, 2, 393, 395, 5, 14, 8, 2, 394, 393,
	3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395, 398, 3, 2, 2, 2, 396, 398, 5, 84,
	43, 2, 397, 383, 3, 2, 2, 2, 397, 390, 3, 2, 2, 2, 397, 396, 3, 2, 2, 2,
	398, 53, 3, 2, 2, 2, 399, 401, 5, 14, 8, 2, 400, 399, 3, 2, 2, 2, 400,
	401, 3, 2, 2, 2, 401, 402, 3, 2, 2, 2, 402, 404, 5, 20, 11, 2, 403, 405,
	5, 14, 8, 2, 404, 403, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 415, 3, 2,
	2, 2, 406, 408, 5, 14, 8, 2, 407, 406, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2,
	408, 409, 3, 2, 2, 2, 409, 411, 5, 56, 29, 2, 410, 412, 5, 14, 8, 2, 411,
	410, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 415, 3, 2, 2, 2, 413, 415,
	5, 86, 44, 2, 414, 400, 3, 2, 2, 2, 414, 407, 3, 2, 2, 2, 414, 413, 3,
	2, 2, 2, 415, 55, 3, 2, 2, 2, 416, 423, 7, 36, 2, 2, 417, 419, 5, 6, 4,
	2, 418, 417, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420,
	422, 5, 58, 30, 2, 421, 418, 3, 2, 2, 2, 422, 425, 3, 2, 2, 2, 423, 421,
	3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 427, 3, 2, 2, 2, 425, 423, 3, 2,
	2, 2, 426, 428, 5, 6, 4, 2, 427, 426, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2,
	428, 429, 3, 2, 2, 2, 429, 430, 7, 38, 2, 2, 430, 57, 3, 2, 2, 2, 431,
	432, 9, 3, 2, 2, 432, 59, 3, 2, 2, 2, 433, 434, 9, 4, 2, 2, 434, 61, 3,
	2, 2, 2, 435, 441, 5, 30, 16, 2, 436, 440, 5, 30, 16, 2, 437, 440, 7, 25,
	2, 2, 438, 440, 5, 14, 8, 2, 439, 436, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2,
	439, 438, 3, 2, 2, 2, 440, 443, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 441,
	442, 3, 2, 2, 2, 442, 63, 3, 2, 2, 2, 443, 441, 3, 2, 2, 2, 444, 445, 5,
	60, 31, 2, 445, 65, 3, 2, 2, 2, 446, 447, 5, 60, 31, 2, 447, 67, 3, 2,
	2, 2, 448, 453, 7, 37, 2, 2, 449, 454, 7, 3, 2, 2, 450, 454, 5, 60, 31,
	2, 451, 454, 7, 6, 2, 2, 452, 454, 7, 9, 2, 2, 453, 449, 3, 2, 2, 2, 453,
	450, 3, 2, 2, 2, 453, 451, 3, 2, 2, 2, 453, 452, 3, 2, 2, 2, 454, 69, 3,
	2, 2, 2, 455, 457, 5, 104, 53, 2, 456, 455, 3, 2, 2, 2, 457, 458, 3, 2,
	2, 2, 458, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2,
	460, 462, 5, 102, 52, 2, 461, 463, 5, 104, 53, 2, 462, 461, 3, 2, 2, 2,
	463, 464, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 464, 465, 3, 2, 2, 2, 465,
	71, 3, 2, 2, 2, 466, 468, 5, 14, 8, 2, 467, 466, 3, 2, 2, 2, 467, 468,
	3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 470, 7, 30, 2, 2, 470, 471, 5, 74,
	38, 2, 471, 472, 5, 50, 26, 2, 472, 474, 7, 32, 2, 2, 473, 475, 5, 14,
	8, 2, 474, 473, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475, 73, 3, 2, 2, 2,
	476, 477, 5, 76, 39, 2, 477, 478, 7, 28, 2, 2, 478, 75, 3, 2, 2, 2, 479,
	482, 5, 14, 8, 2, 480, 482, 7, 23, 2, 2, 481, 479, 3, 2, 2, 2, 481, 480,
	3, 2, 2, 2, 482, 485, 3, 2, 2, 2, 483, 481, 3, 2, 2, 2, 483, 484, 3, 2,
	2, 2, 484, 486, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2, 486, 487, 7, 34, 2, 2,
	487, 498, 5, 54, 28, 2, 488, 490, 7, 23, 2, 2, 489, 491, 5, 14, 8, 2, 490,
	489, 3, 2, 2, 2, 490, 491, 3, 2, 2, 2, 491, 494, 3, 2, 2, 2, 492, 493,
	7, 34, 2, 2, 493, 495, 5, 54, 28, 2, 494, 492, 3, 2, 2, 2, 494, 495, 3,
	2, 2, 2, 495, 497, 3, 2, 2, 2, 496, 488, 3, 2, 2, 2, 497, 500, 3, 2, 2,
	2, 498, 496, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 77, 3, 2, 2, 2, 500,
	498, 3, 2, 2, 2, 501, 503, 5, 14, 8, 2, 502, 501, 3, 2, 2, 2, 502, 503,
	3, 2, 2, 2, 503, 504, 3, 2, 2, 2, 504, 506, 7, 23, 2, 2, 505, 502, 3, 2,
	2, 2, 506, 509, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2, 507, 508, 3, 2, 2, 2,
	508, 510, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 510, 518, 5, 34, 18, 2, 511,
	514, 7, 23, 2, 2, 512, 515, 5, 34, 18, 2, 513, 515, 5, 14, 8, 2, 514, 512,
	3, 2, 2, 2, 514, 513, 3, 2, 2, 2, 514, 515, 3, 2, 2, 2, 515, 517, 3, 2,
	2, 2, 516, 511, 3, 2, 2, 2, 517, 520, 3, 2, 2, 2, 518, 516, 3, 2, 2, 2,
	518, 519, 3, 2, 2, 2, 519, 79, 3, 2, 2, 2, 520, 518, 3, 2, 2, 2, 521, 523,
	5, 14, 8, 2, 522, 521, 3, 2, 2, 2, 522, 523, 3, 2, 2, 2, 523, 524, 3, 2,
	2, 2, 524, 526, 7, 23, 2, 2, 525, 522, 3, 2, 2, 2, 526, 529, 3, 2, 2, 2,
	527, 525, 3, 2, 2, 2, 527, 528, 3, 2, 2, 2, 528, 530, 3, 2, 2, 2, 529,
	527, 3, 2, 2, 2, 530, 538, 5, 32, 17, 2, 531, 534, 7, 23, 2, 2, 532, 535,
	5, 32, 17, 2, 533, 535, 5, 14, 8, 2, 534, 532, 3, 2, 2, 2, 534, 533, 3,
	2, 2, 2, 534, 535, 3, 2, 2, 2, 535, 537, 3, 2, 2, 2, 536, 531, 3, 2, 2,
	2, 537, 540, 3, 2, 2, 2, 538, 536, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539,
	81, 3, 2, 2, 2, 540, 538, 3, 2, 2, 2, 541, 543, 5, 14, 8, 2, 542, 541,
	3, 2, 2, 2, 542, 543, 3, 2, 2, 2, 543, 544, 3, 2, 2, 2, 544, 546, 7, 23,
	2, 2, 545, 542, 3, 2, 2, 2, 546, 547, 3, 2, 2, 2, 547, 545, 3, 2, 2, 2,
	547, 548, 3, 2, 2, 2, 548, 550, 3, 2, 2, 2, 549, 551, 5, 14, 8, 2, 550,
	549, 3, 2, 2, 2, 550, 551, 3, 2, 2, 2, 551, 83, 3, 2, 2, 2, 552, 557, 5,
	30, 16, 2, 553, 554, 7, 25, 2, 2, 554, 556, 5, 30, 16, 2, 555, 553, 3,
	2, 2, 2, 556, 559, 3, 2, 2, 2, 557, 555, 3, 2, 2, 2, 557, 558, 3, 2, 2,
	2, 558, 85, 3, 2, 2, 2, 559, 557, 3, 2, 2, 2, 560, 565, 5, 18, 10, 2, 561,
	562, 7, 25, 2, 2, 562, 564, 5, 18, 10, 2, 563, 561, 3, 2, 2, 2, 564, 567,
	3, 2, 2, 2, 565, 563, 3, 2, 2, 2, 565, 566, 3, 2, 2, 2, 566, 87, 3, 2,
	2, 2, 567, 565, 3, 2, 2, 2, 568, 569, 7, 31, 2, 2, 569, 570, 7, 33, 2,
	2, 570, 571, 5, 90, 46, 2, 571, 572, 7, 33, 2, 2, 572, 573, 5, 92, 47,
	2, 573, 574, 7, 33, 2, 2, 574, 575, 5, 98, 50, 2, 575, 576, 7, 33, 2, 2,
	576, 577, 7, 31, 2, 2, 577, 89, 3, 2, 2, 2, 578, 579, 5, 94, 48, 2, 579,
	91, 3, 2, 2, 2, 580, 581, 5, 94, 48, 2, 581, 93, 3, 2, 2, 2, 582, 584,
	5, 96, 49, 2, 583, 582, 3, 2, 2, 2, 584, 585, 3, 2, 2, 2, 585, 583, 3,
	2, 2, 2, 585, 586, 3, 2, 2, 2, 586, 95, 3, 2, 2, 2, 587, 588, 9, 5, 2,
	2, 588, 97, 3, 2, 2, 2, 589, 591, 5, 100, 51, 2, 590, 589, 3, 2, 2, 2,
	591, 592, 3, 2, 2, 2, 592, 590, 3, 2, 2, 2, 592, 593, 3, 2, 2, 2, 593,
	99, 3, 2, 2, 2, 594, 595, 9, 6, 2, 2, 595, 101, 3, 2, 2, 2, 596, 597, 7,
	9, 2, 2, 597, 598, 7, 6, 2, 2, 598, 103, 3, 2, 2, 2, 599, 600, 9, 7, 2,
	2, 600, 105, 3, 2, 2, 2, 601, 602, 9, 8, 2, 2, 602, 107, 3, 2, 2, 2, 89,
	110, 115, 120, 124, 129, 132, 168, 173, 177, 182, 186, 191, 196, 199, 202,
	209, 214, 220, 224, 262, 266, 269, 274, 280, 285, 289, 292, 296, 299, 303,
	305, 309, 313, 316, 321, 325, 329, 332, 337, 341, 346, 349, 356, 360, 367,
	371, 376, 383, 387, 390, 394, 397, 400, 404, 407, 411, 414, 418, 423, 427,
	439, 441, 453, 458, 464, 467, 474, 481, 483, 490, 494, 498, 502, 507, 514,
	518, 522, 527, 534, 538, 542, 547, 550, 557, 565, 585, 592,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\u0000'", "", "'\t'", "'\n'", "'\u000B'", "'\u000C'", "'\r'", "",
	"' '", "'!'", "'\"'", "'#'", "'$'", "'%'", "'&'", "'''", "'('", "')'",
	"'*'", "'+'", "','", "'-'", "'.'", "'/'", "", "':'", "';'", "'<'", "'='",
	"'>'", "'?'", "'@'", "", "'['", "'\\'", "']'", "'^'", "'_'", "'`'", "",
	"'{'", "'|'", "'}'", "'~'", "'\u007F'",
}
var symbolicNames = []string{
	"", "U_00", "U_01_08", "TAB", "LF", "U_0B", "U_0C", "CR", "U_0E_1F", "SP",
	"Exclamation", "DQuote", "Hash", "Dollar", "Percent", "Ampersand", "SQuote",
	"LParens", "RParens", "Asterisk", "Plus", "Comma", "Minus", "Period", "Slash",
	"Digit", "Colon", "Semicolon", "Less", "Equal", "Greater", "Question",
	"At", "AlphaUpper", "LBracket", "Backslash", "RBracket", "Caret", "Underscore",
	"Backtick", "AlphaLower", "LCurly", "Pipe", "RCurly", "Tilde", "Delete",
	"UTF8NonAscii",
}

var ruleNames = []string{
	"quotedChar", "quotedPair", "fws", "ctext", "ccontent", "comment", "cfws",
	"atext", "atom", "dotAtom", "qtext", "quotedContent", "quotedValue", "quotedString",
	"word", "address", "mailbox", "nameAddr", "angleAddr", "group", "displayName",
	"mailboxList", "addressList", "groupList", "addrSpec", "localPart", "domain",
	"domainLiteral", "dtext", "obsNoWSCTL", "obsPhrase", "obsCtext", "obsQtext",
	"obsQP", "obsFWS", "obsAngleAddr", "obsRoute", "obsDomainList", "obsMboxList",
	"obsAddrList", "obsGroupList", "obsLocalPart", "obsDomain", "encodedWord",
	"charset", "encoding", "token", "tokenChar", "encodedText", "encodedChar",
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
	AddressParserEOF          = antlr.TokenEOF
	AddressParserU_00         = 1
	AddressParserU_01_08      = 2
	AddressParserTAB          = 3
	AddressParserLF           = 4
	AddressParserU_0B         = 5
	AddressParserU_0C         = 6
	AddressParserCR           = 7
	AddressParserU_0E_1F      = 8
	AddressParserSP           = 9
	AddressParserExclamation  = 10
	AddressParserDQuote       = 11
	AddressParserHash         = 12
	AddressParserDollar       = 13
	AddressParserPercent      = 14
	AddressParserAmpersand    = 15
	AddressParserSQuote       = 16
	AddressParserLParens      = 17
	AddressParserRParens      = 18
	AddressParserAsterisk     = 19
	AddressParserPlus         = 20
	AddressParserComma        = 21
	AddressParserMinus        = 22
	AddressParserPeriod       = 23
	AddressParserSlash        = 24
	AddressParserDigit        = 25
	AddressParserColon        = 26
	AddressParserSemicolon    = 27
	AddressParserLess         = 28
	AddressParserEqual        = 29
	AddressParserGreater      = 30
	AddressParserQuestion     = 31
	AddressParserAt           = 32
	AddressParserAlphaUpper   = 33
	AddressParserLBracket     = 34
	AddressParserBackslash    = 35
	AddressParserRBracket     = 36
	AddressParserCaret        = 37
	AddressParserUnderscore   = 38
	AddressParserBacktick     = 39
	AddressParserAlphaLower   = 40
	AddressParserLCurly       = 41
	AddressParserPipe         = 42
	AddressParserRCurly       = 43
	AddressParserTilde        = 44
	AddressParserDelete       = 45
	AddressParserUTF8NonAscii = 46
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
	AddressParserRULE_address       = 15
	AddressParserRULE_mailbox       = 16
	AddressParserRULE_nameAddr      = 17
	AddressParserRULE_angleAddr     = 18
	AddressParserRULE_group         = 19
	AddressParserRULE_displayName   = 20
	AddressParserRULE_mailboxList   = 21
	AddressParserRULE_addressList   = 22
	AddressParserRULE_groupList     = 23
	AddressParserRULE_addrSpec      = 24
	AddressParserRULE_localPart     = 25
	AddressParserRULE_domain        = 26
	AddressParserRULE_domainLiteral = 27
	AddressParserRULE_dtext         = 28
	AddressParserRULE_obsNoWSCTL    = 29
	AddressParserRULE_obsPhrase     = 30
	AddressParserRULE_obsCtext      = 31
	AddressParserRULE_obsQtext      = 32
	AddressParserRULE_obsQP         = 33
	AddressParserRULE_obsFWS        = 34
	AddressParserRULE_obsAngleAddr  = 35
	AddressParserRULE_obsRoute      = 36
	AddressParserRULE_obsDomainList = 37
	AddressParserRULE_obsMboxList   = 38
	AddressParserRULE_obsAddrList   = 39
	AddressParserRULE_obsGroupList  = 40
	AddressParserRULE_obsLocalPart  = 41
	AddressParserRULE_obsDomain     = 42
	AddressParserRULE_encodedWord   = 43
	AddressParserRULE_charset       = 44
	AddressParserRULE_encoding      = 45
	AddressParserRULE_token         = 46
	AddressParserRULE_tokenChar     = 47
	AddressParserRULE_encodedText   = 48
	AddressParserRULE_encodedChar   = 49
	AddressParserRULE_crlf          = 50
	AddressParserRULE_wsp           = 51
	AddressParserRULE_vchar         = 52
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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation, AddressParserDQuote, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserLParens, AddressParserRParens, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserBackslash, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde, AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Vchar()
		}

	case AddressParserTAB, AddressParserSP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
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

func (s *QuotedPairContext) ObsQP() IObsQPContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsQPContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsQPContext)
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

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(AddressParserBackslash)
		}
		{
			p.SetState(111)
			p.QuotedChar()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.ObsQP()
		}

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

func (s *FwsContext) ObsFWS() IObsFWSContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsFWSContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsFWSContext)
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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			p.SetState(118)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AddressParserTAB || _la == AddressParserSP {
				{
					p.SetState(115)
					p.Wsp()
				}

				p.SetState(120)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(121)
				p.Crlf()
			}

		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(124)
					p.Wsp()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.ObsFWS()
		}

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

func (s *CtextContext) ObsCtext() IObsCtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsCtextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsCtextContext)
}

func (s *CtextContext) UTF8NonAscii() antlr.TerminalNode {
	return s.GetToken(AddressParserUTF8NonAscii, 0)
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

	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(AddressParserExclamation)
		}

	case AddressParserDQuote:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(AddressParserDQuote)
		}

	case AddressParserHash:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(AddressParserHash)
		}

	case AddressParserDollar:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)
			p.Match(AddressParserDollar)
		}

	case AddressParserPercent:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(136)
			p.Match(AddressParserPercent)
		}

	case AddressParserAmpersand:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(137)
			p.Match(AddressParserAmpersand)
		}

	case AddressParserSQuote:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(138)
			p.Match(AddressParserSQuote)
		}

	case AddressParserAsterisk:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(139)
			p.Match(AddressParserAsterisk)
		}

	case AddressParserPlus:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(140)
			p.Match(AddressParserPlus)
		}

	case AddressParserComma:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(141)
			p.Match(AddressParserComma)
		}

	case AddressParserMinus:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(142)
			p.Match(AddressParserMinus)
		}

	case AddressParserPeriod:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(143)
			p.Match(AddressParserPeriod)
		}

	case AddressParserSlash:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(144)
			p.Match(AddressParserSlash)
		}

	case AddressParserDigit:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(145)
			p.Match(AddressParserDigit)
		}

	case AddressParserColon:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(146)
			p.Match(AddressParserColon)
		}

	case AddressParserSemicolon:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(147)
			p.Match(AddressParserSemicolon)
		}

	case AddressParserLess:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(148)
			p.Match(AddressParserLess)
		}

	case AddressParserEqual:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(149)
			p.Match(AddressParserEqual)
		}

	case AddressParserGreater:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(150)
			p.Match(AddressParserGreater)
		}

	case AddressParserQuestion:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(151)
			p.Match(AddressParserQuestion)
		}

	case AddressParserAt:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(152)
			p.Match(AddressParserAt)
		}

	case AddressParserAlphaUpper:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(153)
			p.Match(AddressParserAlphaUpper)
		}

	case AddressParserLBracket:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(154)
			p.Match(AddressParserLBracket)
		}

	case AddressParserRBracket:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(155)
			p.Match(AddressParserRBracket)
		}

	case AddressParserCaret:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(156)
			p.Match(AddressParserCaret)
		}

	case AddressParserUnderscore:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(157)
			p.Match(AddressParserUnderscore)
		}

	case AddressParserBacktick:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(158)
			p.Match(AddressParserBacktick)
		}

	case AddressParserAlphaLower:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(159)
			p.Match(AddressParserAlphaLower)
		}

	case AddressParserLCurly:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(160)
			p.Match(AddressParserLCurly)
		}

	case AddressParserPipe:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(161)
			p.Match(AddressParserPipe)
		}

	case AddressParserRCurly:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(162)
			p.Match(AddressParserRCurly)
		}

	case AddressParserTilde:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(163)
			p.Match(AddressParserTilde)
		}

	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserDelete:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(164)
			p.ObsCtext()
		}

	case AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(165)
			p.Match(AddressParserUTF8NonAscii)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserExclamation, AddressParserDQuote, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde, AddressParserDelete, AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Ctext()
		}

	case AddressParserBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)
			p.QuotedPair()
		}

	case AddressParserLParens:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(170)
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
		p.SetState(173)
		p.Match(AddressParserLParens)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(174)
					p.Fws()
				}

			}
			{
				p.SetState(177)
				p.Ccontent()
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(183)
			p.Fws()
		}

	}
	{
		p.SetState(186)
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

	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(189)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
					{
						p.SetState(188)
						p.Fws()
					}

				}
				{
					p.SetState(191)
					p.Comment()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(196)
				p.Fws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
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

func (s *AtextContext) UTF8NonAscii() antlr.TerminalNode {
	return s.GetToken(AddressParserUTF8NonAscii, 0)
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
		p.SetState(202)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33))|(1<<(AddressParserUTF8NonAscii-33)))) != 0)) {
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
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(204)
				p.Atext()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33))|(1<<(AddressParserUTF8NonAscii-33)))) != 0) {
		{
			p.SetState(209)
			p.Atext()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserPeriod {
		{
			p.SetState(214)
			p.Match(AddressParserPeriod)
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33))|(1<<(AddressParserUTF8NonAscii-33)))) != 0) {
			{
				p.SetState(215)
				p.Atext()
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(224)
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

func (s *QtextContext) ObsQtext() IObsQtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsQtextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsQtextContext)
}

func (s *QtextContext) UTF8NonAscii() antlr.TerminalNode {
	return s.GetToken(AddressParserUTF8NonAscii, 0)
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

	p.SetState(260)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.Match(AddressParserExclamation)
		}

	case AddressParserHash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.Match(AddressParserHash)
		}

	case AddressParserDollar:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(227)
			p.Match(AddressParserDollar)
		}

	case AddressParserPercent:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(228)
			p.Match(AddressParserPercent)
		}

	case AddressParserAmpersand:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(229)
			p.Match(AddressParserAmpersand)
		}

	case AddressParserSQuote:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(230)
			p.Match(AddressParserSQuote)
		}

	case AddressParserLParens:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(231)
			p.Match(AddressParserLParens)
		}

	case AddressParserRParens:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(232)
			p.Match(AddressParserRParens)
		}

	case AddressParserAsterisk:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(233)
			p.Match(AddressParserAsterisk)
		}

	case AddressParserPlus:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(234)
			p.Match(AddressParserPlus)
		}

	case AddressParserComma:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(235)
			p.Match(AddressParserComma)
		}

	case AddressParserMinus:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(236)
			p.Match(AddressParserMinus)
		}

	case AddressParserPeriod:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(237)
			p.Match(AddressParserPeriod)
		}

	case AddressParserSlash:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(238)
			p.Match(AddressParserSlash)
		}

	case AddressParserDigit:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(239)
			p.Match(AddressParserDigit)
		}

	case AddressParserColon:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(240)
			p.Match(AddressParserColon)
		}

	case AddressParserSemicolon:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(241)
			p.Match(AddressParserSemicolon)
		}

	case AddressParserLess:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(242)
			p.Match(AddressParserLess)
		}

	case AddressParserEqual:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(243)
			p.Match(AddressParserEqual)
		}

	case AddressParserGreater:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(244)
			p.Match(AddressParserGreater)
		}

	case AddressParserQuestion:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(245)
			p.Match(AddressParserQuestion)
		}

	case AddressParserAt:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(246)
			p.Match(AddressParserAt)
		}

	case AddressParserAlphaUpper:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(247)
			p.Match(AddressParserAlphaUpper)
		}

	case AddressParserLBracket:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(248)
			p.Match(AddressParserLBracket)
		}

	case AddressParserRBracket:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(249)
			p.Match(AddressParserRBracket)
		}

	case AddressParserCaret:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(250)
			p.Match(AddressParserCaret)
		}

	case AddressParserUnderscore:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(251)
			p.Match(AddressParserUnderscore)
		}

	case AddressParserBacktick:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(252)
			p.Match(AddressParserBacktick)
		}

	case AddressParserAlphaLower:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(253)
			p.Match(AddressParserAlphaLower)
		}

	case AddressParserLCurly:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(254)
			p.Match(AddressParserLCurly)
		}

	case AddressParserPipe:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(255)
			p.Match(AddressParserPipe)
		}

	case AddressParserRCurly:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(256)
			p.Match(AddressParserRCurly)
		}

	case AddressParserTilde:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(257)
			p.Match(AddressParserTilde)
		}

	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserDelete:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(258)
			p.ObsQtext()
		}

	case AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(259)
			p.Match(AddressParserUTF8NonAscii)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserExclamation, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserLParens, AddressParserRParens, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde, AddressParserDelete, AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Qtext()
		}

	case AddressParserBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
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
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(266)
					p.Fws()
				}

			}
			{
				p.SetState(269)
				p.QuotedContent()
			}

		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
		p.SetState(275)
		p.Match(AddressParserDQuote)
	}
	{
		p.SetState(276)
		p.QuotedValue()
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(277)
			p.Fws()
		}

	}
	{
		p.SetState(280)
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

	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(282)
				p.Cfws()
			}

		}
		{
			p.SetState(285)
			p.EncodedWord()
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(286)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(289)
				p.Cfws()
			}

		}
		{
			p.SetState(292)
			p.Atom()
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(293)
				p.Cfws()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(296)
				p.Cfws()
			}

		}
		{
			p.SetState(299)
			p.QuotedString()
		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(300)
				p.Cfws()
			}

		}

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
	p.EnterRule(localctx, 30, AddressParserRULE_address)

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

	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.Mailbox()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
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
	p.EnterRule(localctx, 32, AddressParserRULE_mailbox)

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

	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.NameAddr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
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
	p.EnterRule(localctx, 34, AddressParserRULE_nameAddr)

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
	p.SetState(314)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(313)
			p.DisplayName()
		}

	}
	{
		p.SetState(316)
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

func (s *AngleAddrContext) AddrSpec() IAddrSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddrSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddrSpecContext)
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
	p.EnterRule(localctx, 36, AddressParserRULE_angleAddr)
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

	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(318)
				p.Cfws()
			}

		}
		{
			p.SetState(321)
			p.Match(AddressParserLess)
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33))|(1<<(AddressParserUTF8NonAscii-33)))) != 0) {
			{
				p.SetState(322)
				p.AddrSpec()
			}

		}
		{
			p.SetState(325)
			p.Match(AddressParserGreater)
		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(326)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
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
	p.EnterRule(localctx, 38, AddressParserRULE_group)
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
		p.SetState(332)
		p.DisplayName()
	}
	{
		p.SetState(333)
		p.Match(AddressParserColon)
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33))|(1<<(AddressParserUTF8NonAscii-33)))) != 0) {
		{
			p.SetState(334)
			p.GroupList()
		}

	}
	{
		p.SetState(337)
		p.Match(AddressParserSemicolon)
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(338)
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
	p.EnterRule(localctx, 40, AddressParserRULE_displayName)

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

	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(341)
					p.Word()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(344)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
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
	p.EnterRule(localctx, 42, AddressParserRULE_mailboxList)
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

	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.Mailbox()
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AddressParserComma {
			{
				p.SetState(350)
				p.Match(AddressParserComma)
			}
			{
				p.SetState(351)
				p.Mailbox()
			}

			p.SetState(356)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
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
	p.EnterRule(localctx, 44, AddressParserRULE_addressList)
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

	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)
			p.Address()
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AddressParserComma {
			{
				p.SetState(361)
				p.Match(AddressParserComma)
			}
			{
				p.SetState(362)
				p.Address()
			}

			p.SetState(367)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
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
	p.EnterRule(localctx, 46, AddressParserRULE_groupList)

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

	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.MailboxList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.Cfws()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
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
	p.EnterRule(localctx, 48, AddressParserRULE_addrSpec)

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
		p.SetState(376)
		p.LocalPart()
	}
	{
		p.SetState(377)
		p.Match(AddressParserAt)
	}
	{
		p.SetState(378)
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

func (s *LocalPartContext) ObsLocalPart() IObsLocalPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsLocalPartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsLocalPartContext)
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
	p.EnterRule(localctx, 50, AddressParserRULE_localPart)
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

	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(380)
				p.Cfws()
			}

		}
		{
			p.SetState(383)
			p.DotAtom()
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(384)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(387)
				p.Cfws()
			}

		}
		{
			p.SetState(390)
			p.QuotedString()
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(391)
				p.Cfws()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(394)
			p.ObsLocalPart()
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

func (s *DomainContext) ObsDomain() IObsDomainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsDomainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsDomainContext)
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
	p.EnterRule(localctx, 52, AddressParserRULE_domain)
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

	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(397)
				p.Cfws()
			}

		}
		{
			p.SetState(400)
			p.DotAtom()
		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(401)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(404)
				p.Cfws()
			}

		}
		{
			p.SetState(407)
			p.DomainLiteral()
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(408)
				p.Cfws()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
			p.ObsDomain()
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
	p.EnterRule(localctx, 54, AddressParserRULE_domainLiteral)
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
		p.SetState(414)
		p.Match(AddressParserLBracket)
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(416)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(415)
					p.Fws()
				}

			}
			{
				p.SetState(418)
				p.Dtext()
			}

		}
		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(424)
			p.Fws()
		}

	}
	{
		p.SetState(427)
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

func (s *DtextContext) UTF8NonAscii() antlr.TerminalNode {
	return s.GetToken(AddressParserUTF8NonAscii, 0)
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
	p.EnterRule(localctx, 56, AddressParserRULE_dtext)
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
		p.SetState(429)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater)|(1<<AddressParserQuestion))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserAt-32))|(1<<(AddressParserAlphaUpper-32))|(1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32))|(1<<(AddressParserUTF8NonAscii-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IObsNoWSCTLContext is an interface to support dynamic dispatch.
type IObsNoWSCTLContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsNoWSCTLContext differentiates from other interfaces.
	IsObsNoWSCTLContext()
}

type ObsNoWSCTLContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsNoWSCTLContext() *ObsNoWSCTLContext {
	var p = new(ObsNoWSCTLContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsNoWSCTL
	return p
}

func (*ObsNoWSCTLContext) IsObsNoWSCTLContext() {}

func NewObsNoWSCTLContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsNoWSCTLContext {
	var p = new(ObsNoWSCTLContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsNoWSCTL

	return p
}

func (s *ObsNoWSCTLContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsNoWSCTLContext) U_01_08() antlr.TerminalNode {
	return s.GetToken(AddressParserU_01_08, 0)
}

func (s *ObsNoWSCTLContext) U_0B() antlr.TerminalNode {
	return s.GetToken(AddressParserU_0B, 0)
}

func (s *ObsNoWSCTLContext) U_0C() antlr.TerminalNode {
	return s.GetToken(AddressParserU_0C, 0)
}

func (s *ObsNoWSCTLContext) U_0E_1F() antlr.TerminalNode {
	return s.GetToken(AddressParserU_0E_1F, 0)
}

func (s *ObsNoWSCTLContext) Delete() antlr.TerminalNode {
	return s.GetToken(AddressParserDelete, 0)
}

func (s *ObsNoWSCTLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsNoWSCTLContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsNoWSCTLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsNoWSCTL(s)
	}
}

func (s *ObsNoWSCTLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsNoWSCTL(s)
	}
}

func (p *AddressParser) ObsNoWSCTL() (localctx IObsNoWSCTLContext) {
	localctx = NewObsNoWSCTLContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, AddressParserRULE_obsNoWSCTL)
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
		p.SetState(431)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserU_01_08)|(1<<AddressParserU_0B)|(1<<AddressParserU_0C)|(1<<AddressParserU_0E_1F))) != 0) || _la == AddressParserDelete) {
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
		p.SetState(433)
		p.Word()
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(437)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(434)
					p.Word()
				}

			case 2:
				{
					p.SetState(435)
					p.Match(AddressParserPeriod)
				}

			case 3:
				{
					p.SetState(436)
					p.Cfws()
				}

			}

		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())
	}

	return localctx
}

// IObsCtextContext is an interface to support dynamic dispatch.
type IObsCtextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsCtextContext differentiates from other interfaces.
	IsObsCtextContext()
}

type ObsCtextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsCtextContext() *ObsCtextContext {
	var p = new(ObsCtextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsCtext
	return p
}

func (*ObsCtextContext) IsObsCtextContext() {}

func NewObsCtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsCtextContext {
	var p = new(ObsCtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsCtext

	return p
}

func (s *ObsCtextContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsCtextContext) ObsNoWSCTL() IObsNoWSCTLContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsNoWSCTLContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsNoWSCTLContext)
}

func (s *ObsCtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsCtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsCtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsCtext(s)
	}
}

func (s *ObsCtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsCtext(s)
	}
}

func (p *AddressParser) ObsCtext() (localctx IObsCtextContext) {
	localctx = NewObsCtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AddressParserRULE_obsCtext)

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
		p.SetState(442)
		p.ObsNoWSCTL()
	}

	return localctx
}

// IObsQtextContext is an interface to support dynamic dispatch.
type IObsQtextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsQtextContext differentiates from other interfaces.
	IsObsQtextContext()
}

type ObsQtextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsQtextContext() *ObsQtextContext {
	var p = new(ObsQtextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsQtext
	return p
}

func (*ObsQtextContext) IsObsQtextContext() {}

func NewObsQtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsQtextContext {
	var p = new(ObsQtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsQtext

	return p
}

func (s *ObsQtextContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsQtextContext) ObsNoWSCTL() IObsNoWSCTLContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsNoWSCTLContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsNoWSCTLContext)
}

func (s *ObsQtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsQtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsQtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsQtext(s)
	}
}

func (s *ObsQtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsQtext(s)
	}
}

func (p *AddressParser) ObsQtext() (localctx IObsQtextContext) {
	localctx = NewObsQtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AddressParserRULE_obsQtext)

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
		p.SetState(444)
		p.ObsNoWSCTL()
	}

	return localctx
}

// IObsQPContext is an interface to support dynamic dispatch.
type IObsQPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsQPContext differentiates from other interfaces.
	IsObsQPContext()
}

type ObsQPContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsQPContext() *ObsQPContext {
	var p = new(ObsQPContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsQP
	return p
}

func (*ObsQPContext) IsObsQPContext() {}

func NewObsQPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsQPContext {
	var p = new(ObsQPContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsQP

	return p
}

func (s *ObsQPContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsQPContext) Backslash() antlr.TerminalNode {
	return s.GetToken(AddressParserBackslash, 0)
}

func (s *ObsQPContext) U_00() antlr.TerminalNode {
	return s.GetToken(AddressParserU_00, 0)
}

func (s *ObsQPContext) ObsNoWSCTL() IObsNoWSCTLContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObsNoWSCTLContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObsNoWSCTLContext)
}

func (s *ObsQPContext) LF() antlr.TerminalNode {
	return s.GetToken(AddressParserLF, 0)
}

func (s *ObsQPContext) CR() antlr.TerminalNode {
	return s.GetToken(AddressParserCR, 0)
}

func (s *ObsQPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsQPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsQPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsQP(s)
	}
}

func (s *ObsQPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsQP(s)
	}
}

func (p *AddressParser) ObsQP() (localctx IObsQPContext) {
	localctx = NewObsQPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, AddressParserRULE_obsQP)

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
		p.SetState(446)
		p.Match(AddressParserBackslash)
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserU_00:
		{
			p.SetState(447)
			p.Match(AddressParserU_00)
		}

	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserDelete:
		{
			p.SetState(448)
			p.ObsNoWSCTL()
		}

	case AddressParserLF:
		{
			p.SetState(449)
			p.Match(AddressParserLF)
		}

	case AddressParserCR:
		{
			p.SetState(450)
			p.Match(AddressParserCR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObsFWSContext is an interface to support dynamic dispatch.
type IObsFWSContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsFWSContext differentiates from other interfaces.
	IsObsFWSContext()
}

type ObsFWSContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsFWSContext() *ObsFWSContext {
	var p = new(ObsFWSContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsFWS
	return p
}

func (*ObsFWSContext) IsObsFWSContext() {}

func NewObsFWSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsFWSContext {
	var p = new(ObsFWSContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsFWS

	return p
}

func (s *ObsFWSContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsFWSContext) Crlf() ICrlfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICrlfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICrlfContext)
}

func (s *ObsFWSContext) AllWsp() []IWspContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWspContext)(nil)).Elem())
	var tst = make([]IWspContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWspContext)
		}
	}

	return tst
}

func (s *ObsFWSContext) Wsp(i int) IWspContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWspContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWspContext)
}

func (s *ObsFWSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsFWSContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsFWSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsFWS(s)
	}
}

func (s *ObsFWSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsFWS(s)
	}
}

func (p *AddressParser) ObsFWS() (localctx IObsFWSContext) {
	localctx = NewObsFWSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, AddressParserRULE_obsFWS)
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
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AddressParserTAB || _la == AddressParserSP {
		{
			p.SetState(453)
			p.Wsp()
		}

		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(458)
		p.Crlf()
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(459)
				p.Wsp()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 70, AddressParserRULE_obsAngleAddr)
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
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(464)
			p.Cfws()
		}

	}
	{
		p.SetState(467)
		p.Match(AddressParserLess)
	}
	{
		p.SetState(468)
		p.ObsRoute()
	}
	{
		p.SetState(469)
		p.AddrSpec()
	}
	{
		p.SetState(470)
		p.Match(AddressParserGreater)
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(471)
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
	p.EnterRule(localctx, 72, AddressParserRULE_obsRoute)

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
		p.SetState(474)
		p.ObsDomainList()
	}
	{
		p.SetState(475)
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
	p.EnterRule(localctx, 74, AddressParserRULE_obsDomainList)
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
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens)|(1<<AddressParserComma))) != 0 {
		p.SetState(479)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AddressParserTAB, AddressParserCR, AddressParserSP, AddressParserLParens:
			{
				p.SetState(477)
				p.Cfws()
			}

		case AddressParserComma:
			{
				p.SetState(478)
				p.Match(AddressParserComma)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(484)
		p.Match(AddressParserAt)
	}
	{
		p.SetState(485)
		p.Domain()
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(486)
			p.Match(AddressParserComma)
		}
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(487)
				p.Cfws()
			}

		}
		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AddressParserAt {
			{
				p.SetState(490)
				p.Match(AddressParserAt)
			}
			{
				p.SetState(491)
				p.Domain()
			}

		}

		p.SetState(498)
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
	p.EnterRule(localctx, 76, AddressParserRULE_obsMboxList)
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
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(500)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(499)
					p.Cfws()
				}

			}
			{
				p.SetState(502)
				p.Match(AddressParserComma)
			}

		}
		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext())
	}
	{
		p.SetState(508)
		p.Mailbox()
	}
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(509)
			p.Match(AddressParserComma)
		}
		p.SetState(512)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(510)
				p.Mailbox()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(511)
				p.Cfws()
			}

		}

		p.SetState(518)
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
	p.EnterRule(localctx, 78, AddressParserRULE_obsAddrList)
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
	p.SetState(525)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 77, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(520)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(519)
					p.Cfws()
				}

			}
			{
				p.SetState(522)
				p.Match(AddressParserComma)
			}

		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 77, p.GetParserRuleContext())
	}
	{
		p.SetState(528)
		p.Address()
	}
	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(529)
			p.Match(AddressParserComma)
		}
		p.SetState(532)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(530)
				p.Address()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(531)
				p.Cfws()
			}

		}

		p.SetState(538)
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
	p.EnterRule(localctx, 80, AddressParserRULE_obsGroupList)
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
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(540)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(539)
					p.Cfws()
				}

			}
			{
				p.SetState(542)
				p.Match(AddressParserComma)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(545)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 81, p.GetParserRuleContext())
	}
	p.SetState(548)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(547)
			p.Cfws()
		}

	}

	return localctx
}

// IObsLocalPartContext is an interface to support dynamic dispatch.
type IObsLocalPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsLocalPartContext differentiates from other interfaces.
	IsObsLocalPartContext()
}

type ObsLocalPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsLocalPartContext() *ObsLocalPartContext {
	var p = new(ObsLocalPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsLocalPart
	return p
}

func (*ObsLocalPartContext) IsObsLocalPartContext() {}

func NewObsLocalPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsLocalPartContext {
	var p = new(ObsLocalPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsLocalPart

	return p
}

func (s *ObsLocalPartContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsLocalPartContext) AllWord() []IWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordContext)(nil)).Elem())
	var tst = make([]IWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordContext)
		}
	}

	return tst
}

func (s *ObsLocalPartContext) Word(i int) IWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *ObsLocalPartContext) AllPeriod() []antlr.TerminalNode {
	return s.GetTokens(AddressParserPeriod)
}

func (s *ObsLocalPartContext) Period(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, i)
}

func (s *ObsLocalPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsLocalPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsLocalPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsLocalPart(s)
	}
}

func (s *ObsLocalPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsLocalPart(s)
	}
}

func (p *AddressParser) ObsLocalPart() (localctx IObsLocalPartContext) {
	localctx = NewObsLocalPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, AddressParserRULE_obsLocalPart)
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
		p.SetState(550)
		p.Word()
	}
	p.SetState(555)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserPeriod {
		{
			p.SetState(551)
			p.Match(AddressParserPeriod)
		}
		{
			p.SetState(552)
			p.Word()
		}

		p.SetState(557)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObsDomainContext is an interface to support dynamic dispatch.
type IObsDomainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObsDomainContext differentiates from other interfaces.
	IsObsDomainContext()
}

type ObsDomainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObsDomainContext() *ObsDomainContext {
	var p = new(ObsDomainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_obsDomain
	return p
}

func (*ObsDomainContext) IsObsDomainContext() {}

func NewObsDomainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObsDomainContext {
	var p = new(ObsDomainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_obsDomain

	return p
}

func (s *ObsDomainContext) GetParser() antlr.Parser { return s.parser }

func (s *ObsDomainContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *ObsDomainContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ObsDomainContext) AllPeriod() []antlr.TerminalNode {
	return s.GetTokens(AddressParserPeriod)
}

func (s *ObsDomainContext) Period(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserPeriod, i)
}

func (s *ObsDomainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObsDomainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObsDomainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterObsDomain(s)
	}
}

func (s *ObsDomainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitObsDomain(s)
	}
}

func (p *AddressParser) ObsDomain() (localctx IObsDomainContext) {
	localctx = NewObsDomainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, AddressParserRULE_obsDomain)
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
		p.SetState(558)
		p.Atom()
	}
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserPeriod {
		{
			p.SetState(559)
			p.Match(AddressParserPeriod)
		}
		{
			p.SetState(560)
			p.Atom()
		}

		p.SetState(565)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 86, AddressParserRULE_encodedWord)

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
		p.SetState(566)
		p.Match(AddressParserEqual)
	}
	{
		p.SetState(567)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(568)
		p.Charset()
	}
	{
		p.SetState(569)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(570)
		p.Encoding()
	}
	{
		p.SetState(571)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(572)
		p.EncodedText()
	}
	{
		p.SetState(573)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(574)
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
	p.EnterRule(localctx, 88, AddressParserRULE_charset)

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
		p.SetState(576)
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
	p.EnterRule(localctx, 90, AddressParserRULE_encoding)

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
		p.SetState(578)
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
	p.EnterRule(localctx, 92, AddressParserRULE_token)
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
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserDigit))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserBackslash-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33)))) != 0) {
		{
			p.SetState(580)
			p.TokenChar()
		}

		p.SetState(583)
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
	p.EnterRule(localctx, 94, AddressParserRULE_tokenChar)
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
		p.SetState(585)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserDigit))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserBackslash-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33)))) != 0)) {
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
	p.EnterRule(localctx, 96, AddressParserRULE_encodedText)
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
	p.SetState(588)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserAt-32))|(1<<(AddressParserAlphaUpper-32))|(1<<(AddressParserLBracket-32))|(1<<(AddressParserBackslash-32))|(1<<(AddressParserRBracket-32))|(1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0) {
		{
			p.SetState(587)
			p.EncodedChar()
		}

		p.SetState(590)
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
	p.EnterRule(localctx, 98, AddressParserRULE_encodedChar)
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
		p.SetState(592)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserAt-32))|(1<<(AddressParserAlphaUpper-32))|(1<<(AddressParserLBracket-32))|(1<<(AddressParserBackslash-32))|(1<<(AddressParserRBracket-32))|(1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0)) {
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
	p.EnterRule(localctx, 100, AddressParserRULE_crlf)

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
		p.SetState(594)
		p.Match(AddressParserCR)
	}
	{
		p.SetState(595)
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
	p.EnterRule(localctx, 102, AddressParserRULE_wsp)
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
		p.SetState(597)
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

func (s *VcharContext) UTF8NonAscii() antlr.TerminalNode {
	return s.GetToken(AddressParserUTF8NonAscii, 0)
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
	p.EnterRule(localctx, 104, AddressParserRULE_vchar)
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
		p.SetState(599)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater)|(1<<AddressParserQuestion))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserAt-32))|(1<<(AddressParserAlphaUpper-32))|(1<<(AddressParserLBracket-32))|(1<<(AddressParserBackslash-32))|(1<<(AddressParserRBracket-32))|(1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32))|(1<<(AddressParserUTF8NonAscii-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
