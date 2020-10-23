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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 628,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 3, 2, 3, 2, 5, 2, 113, 10, 2, 3, 3, 3, 3, 3, 3, 5, 3, 118, 10, 3,
	3, 4, 7, 4, 121, 10, 4, 12, 4, 14, 4, 124, 11, 4, 3, 4, 5, 4, 127, 10,
	4, 3, 4, 6, 4, 130, 10, 4, 13, 4, 14, 4, 131, 3, 4, 5, 4, 135, 10, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 171, 10,
	5, 3, 6, 3, 6, 3, 6, 5, 6, 176, 10, 6, 3, 7, 3, 7, 5, 7, 180, 10, 7, 3,
	7, 7, 7, 183, 10, 7, 12, 7, 14, 7, 186, 11, 7, 3, 7, 5, 7, 189, 10, 7,
	3, 7, 3, 7, 3, 8, 5, 8, 194, 10, 8, 3, 8, 6, 8, 197, 10, 8, 13, 8, 14,
	8, 198, 3, 8, 5, 8, 202, 10, 8, 3, 8, 5, 8, 205, 10, 8, 3, 9, 3, 9, 3,
	10, 6, 10, 210, 10, 10, 13, 10, 14, 10, 211, 3, 11, 6, 11, 215, 10, 11,
	13, 11, 14, 11, 216, 3, 11, 3, 11, 6, 11, 221, 10, 11, 13, 11, 14, 11,
	222, 7, 11, 225, 10, 11, 12, 11, 14, 11, 228, 11, 11, 3, 11, 5, 11, 231,
	10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 268, 10, 12, 3, 13, 3, 13, 5,
	13, 272, 10, 13, 3, 14, 5, 14, 275, 10, 14, 3, 14, 7, 14, 278, 10, 14,
	12, 14, 14, 14, 281, 11, 14, 3, 15, 3, 15, 3, 15, 5, 15, 286, 10, 15, 3,
	15, 3, 15, 3, 16, 5, 16, 291, 10, 16, 3, 16, 3, 16, 5, 16, 295, 10, 16,
	3, 16, 5, 16, 298, 10, 16, 3, 16, 3, 16, 5, 16, 302, 10, 16, 3, 16, 5,
	16, 305, 10, 16, 3, 16, 3, 16, 5, 16, 309, 10, 16, 3, 16, 5, 16, 312, 10,
	16, 3, 17, 3, 17, 5, 17, 316, 10, 17, 3, 18, 3, 18, 5, 18, 320, 10, 18,
	3, 19, 5, 19, 323, 10, 19, 3, 19, 3, 19, 3, 20, 5, 20, 328, 10, 20, 3,
	20, 3, 20, 5, 20, 332, 10, 20, 3, 20, 3, 20, 5, 20, 336, 10, 20, 3, 20,
	5, 20, 339, 10, 20, 3, 21, 3, 21, 3, 21, 5, 21, 344, 10, 21, 3, 21, 3,
	21, 5, 21, 348, 10, 21, 3, 22, 6, 22, 351, 10, 22, 13, 22, 14, 22, 352,
	3, 22, 5, 22, 356, 10, 22, 3, 23, 3, 23, 3, 23, 7, 23, 361, 10, 23, 12,
	23, 14, 23, 364, 11, 23, 3, 23, 5, 23, 367, 10, 23, 3, 24, 3, 24, 3, 24,
	7, 24, 372, 10, 24, 12, 24, 14, 24, 375, 11, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 5, 24, 382, 10, 24, 3, 25, 3, 25, 3, 25, 5, 25, 387, 10, 25,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 394, 10, 26, 3, 27, 5, 27, 397,
	10, 27, 3, 27, 3, 27, 5, 27, 401, 10, 27, 3, 27, 5, 27, 404, 10, 27, 3,
	27, 3, 27, 5, 27, 408, 10, 27, 3, 27, 5, 27, 411, 10, 27, 3, 28, 6, 28,
	414, 10, 28, 13, 28, 14, 28, 415, 3, 29, 5, 29, 419, 10, 29, 3, 29, 3,
	29, 5, 29, 423, 10, 29, 3, 29, 5, 29, 426, 10, 29, 3, 29, 3, 29, 5, 29,
	430, 10, 29, 3, 29, 5, 29, 433, 10, 29, 3, 29, 3, 29, 5, 29, 437, 10, 29,
	5, 29, 439, 10, 29, 3, 30, 3, 30, 5, 30, 443, 10, 30, 3, 30, 7, 30, 446,
	10, 30, 12, 30, 14, 30, 449, 11, 30, 3, 30, 5, 30, 452, 10, 30, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 464,
	10, 33, 12, 33, 14, 33, 467, 11, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 478, 10, 36, 3, 37, 6, 37, 481, 10,
	37, 13, 37, 14, 37, 482, 3, 37, 3, 37, 6, 37, 487, 10, 37, 13, 37, 14,
	37, 488, 3, 38, 5, 38, 492, 10, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	5, 38, 499, 10, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 7, 40, 506, 10,
	40, 12, 40, 14, 40, 509, 11, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 515,
	10, 40, 3, 40, 3, 40, 5, 40, 519, 10, 40, 7, 40, 521, 10, 40, 12, 40, 14,
	40, 524, 11, 40, 3, 41, 5, 41, 527, 10, 41, 3, 41, 7, 41, 530, 10, 41,
	12, 41, 14, 41, 533, 11, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 539, 10,
	41, 7, 41, 541, 10, 41, 12, 41, 14, 41, 544, 11, 41, 3, 42, 5, 42, 547,
	10, 42, 3, 42, 7, 42, 550, 10, 42, 12, 42, 14, 42, 553, 11, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 5, 42, 559, 10, 42, 7, 42, 561, 10, 42, 12, 42, 14,
	42, 564, 11, 42, 3, 43, 5, 43, 567, 10, 43, 3, 43, 6, 43, 570, 10, 43,
	13, 43, 14, 43, 571, 3, 43, 5, 43, 575, 10, 43, 3, 44, 3, 44, 3, 44, 7,
	44, 580, 10, 44, 12, 44, 14, 44, 583, 11, 44, 3, 45, 3, 45, 3, 45, 7, 45,
	588, 10, 45, 12, 45, 14, 45, 591, 11, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49,
	6, 49, 608, 10, 49, 13, 49, 14, 49, 609, 3, 50, 3, 50, 3, 51, 6, 51, 615,
	10, 51, 13, 51, 14, 51, 616, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 54,
	3, 54, 3, 55, 3, 55, 3, 55, 2, 2, 56, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92,
	94, 96, 98, 100, 102, 104, 106, 108, 2, 9, 12, 2, 12, 12, 14, 18, 21, 22,
	24, 24, 26, 27, 31, 31, 33, 33, 35, 35, 39, 46, 48, 48, 5, 2, 12, 35, 39,
	46, 48, 48, 6, 2, 4, 4, 7, 8, 10, 10, 47, 47, 10, 2, 12, 12, 14, 18, 21,
	22, 24, 24, 27, 27, 35, 35, 37, 37, 39, 46, 4, 2, 12, 32, 34, 46, 4, 2,
	5, 5, 11, 11, 4, 2, 12, 46, 48, 48, 2, 741, 2, 112, 3, 2, 2, 2, 4, 117,
	3, 2, 2, 2, 6, 134, 3, 2, 2, 2, 8, 170, 3, 2, 2, 2, 10, 175, 3, 2, 2, 2,
	12, 177, 3, 2, 2, 2, 14, 204, 3, 2, 2, 2, 16, 206, 3, 2, 2, 2, 18, 209,
	3, 2, 2, 2, 20, 214, 3, 2, 2, 2, 22, 267, 3, 2, 2, 2, 24, 271, 3, 2, 2,
	2, 26, 279, 3, 2, 2, 2, 28, 282, 3, 2, 2, 2, 30, 311, 3, 2, 2, 2, 32, 315,
	3, 2, 2, 2, 34, 319, 3, 2, 2, 2, 36, 322, 3, 2, 2, 2, 38, 338, 3, 2, 2,
	2, 40, 340, 3, 2, 2, 2, 42, 355, 3, 2, 2, 2, 44, 366, 3, 2, 2, 2, 46, 381,
	3, 2, 2, 2, 48, 386, 3, 2, 2, 2, 50, 388, 3, 2, 2, 2, 52, 410, 3, 2, 2,
	2, 54, 413, 3, 2, 2, 2, 56, 438, 3, 2, 2, 2, 58, 440, 3, 2, 2, 2, 60, 455,
	3, 2, 2, 2, 62, 457, 3, 2, 2, 2, 64, 459, 3, 2, 2, 2, 66, 468, 3, 2, 2,
	2, 68, 470, 3, 2, 2, 2, 70, 472, 3, 2, 2, 2, 72, 480, 3, 2, 2, 2, 74, 491,
	3, 2, 2, 2, 76, 500, 3, 2, 2, 2, 78, 507, 3, 2, 2, 2, 80, 531, 3, 2, 2,
	2, 82, 551, 3, 2, 2, 2, 84, 569, 3, 2, 2, 2, 86, 576, 3, 2, 2, 2, 88, 584,
	3, 2, 2, 2, 90, 592, 3, 2, 2, 2, 92, 602, 3, 2, 2, 2, 94, 604, 3, 2, 2,
	2, 96, 607, 3, 2, 2, 2, 98, 611, 3, 2, 2, 2, 100, 614, 3, 2, 2, 2, 102,
	618, 3, 2, 2, 2, 104, 620, 3, 2, 2, 2, 106, 623, 3, 2, 2, 2, 108, 625,
	3, 2, 2, 2, 110, 113, 5, 108, 55, 2, 111, 113, 5, 106, 54, 2, 112, 110,
	3, 2, 2, 2, 112, 111, 3, 2, 2, 2, 113, 3, 3, 2, 2, 2, 114, 115, 7, 37,
	2, 2, 115, 118, 5, 2, 2, 2, 116, 118, 5, 70, 36, 2, 117, 114, 3, 2, 2,
	2, 117, 116, 3, 2, 2, 2, 118, 5, 3, 2, 2, 2, 119, 121, 5, 106, 54, 2, 120,
	119, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123,
	3, 2, 2, 2, 123, 125, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 127, 5, 104,
	53, 2, 126, 122, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2,
	128, 130, 5, 106, 54, 2, 129, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131,
	129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 135,
	5, 72, 37, 2, 134, 126, 3, 2, 2, 2, 134, 133, 3, 2, 2, 2, 135, 7, 3, 2,
	2, 2, 136, 171, 7, 12, 2, 2, 137, 171, 7, 13, 2, 2, 138, 171, 7, 14, 2,
	2, 139, 171, 7, 15, 2, 2, 140, 171, 7, 16, 2, 2, 141, 171, 7, 17, 2, 2,
	142, 171, 7, 18, 2, 2, 143, 171, 7, 21, 2, 2, 144, 171, 7, 22, 2, 2, 145,
	171, 7, 23, 2, 2, 146, 171, 7, 24, 2, 2, 147, 171, 7, 25, 2, 2, 148, 171,
	7, 26, 2, 2, 149, 171, 7, 27, 2, 2, 150, 171, 7, 28, 2, 2, 151, 171, 7,
	29, 2, 2, 152, 171, 7, 30, 2, 2, 153, 171, 7, 31, 2, 2, 154, 171, 7, 32,
	2, 2, 155, 171, 7, 33, 2, 2, 156, 171, 7, 34, 2, 2, 157, 171, 7, 35, 2,
	2, 158, 171, 7, 36, 2, 2, 159, 171, 7, 38, 2, 2, 160, 171, 7, 39, 2, 2,
	161, 171, 7, 40, 2, 2, 162, 171, 7, 41, 2, 2, 163, 171, 7, 42, 2, 2, 164,
	171, 7, 43, 2, 2, 165, 171, 7, 44, 2, 2, 166, 171, 7, 45, 2, 2, 167, 171,
	7, 46, 2, 2, 168, 171, 5, 66, 34, 2, 169, 171, 7, 48, 2, 2, 170, 136, 3,
	2, 2, 2, 170, 137, 3, 2, 2, 2, 170, 138, 3, 2, 2, 2, 170, 139, 3, 2, 2,
	2, 170, 140, 3, 2, 2, 2, 170, 141, 3, 2, 2, 2, 170, 142, 3, 2, 2, 2, 170,
	143, 3, 2, 2, 2, 170, 144, 3, 2, 2, 2, 170, 145, 3, 2, 2, 2, 170, 146,
	3, 2, 2, 2, 170, 147, 3, 2, 2, 2, 170, 148, 3, 2, 2, 2, 170, 149, 3, 2,
	2, 2, 170, 150, 3, 2, 2, 2, 170, 151, 3, 2, 2, 2, 170, 152, 3, 2, 2, 2,
	170, 153, 3, 2, 2, 2, 170, 154, 3, 2, 2, 2, 170, 155, 3, 2, 2, 2, 170,
	156, 3, 2, 2, 2, 170, 157, 3, 2, 2, 2, 170, 158, 3, 2, 2, 2, 170, 159,
	3, 2, 2, 2, 170, 160, 3, 2, 2, 2, 170, 161, 3, 2, 2, 2, 170, 162, 3, 2,
	2, 2, 170, 163, 3, 2, 2, 2, 170, 164, 3, 2, 2, 2, 170, 165, 3, 2, 2, 2,
	170, 166, 3, 2, 2, 2, 170, 167, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170,
	169, 3, 2, 2, 2, 171, 9, 3, 2, 2, 2, 172, 176, 5, 8, 5, 2, 173, 176, 5,
	4, 3, 2, 174, 176, 5, 12, 7, 2, 175, 172, 3, 2, 2, 2, 175, 173, 3, 2, 2,
	2, 175, 174, 3, 2, 2, 2, 176, 11, 3, 2, 2, 2, 177, 184, 7, 19, 2, 2, 178,
	180, 5, 6, 4, 2, 179, 178, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181,
	3, 2, 2, 2, 181, 183, 5, 10, 6, 2, 182, 179, 3, 2, 2, 2, 183, 186, 3, 2,
	2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2,
	186, 184, 3, 2, 2, 2, 187, 189, 5, 6, 4, 2, 188, 187, 3, 2, 2, 2, 188,
	189, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 7, 20, 2, 2, 191, 13,
	3, 2, 2, 2, 192, 194, 5, 6, 4, 2, 193, 192, 3, 2, 2, 2, 193, 194, 3, 2,
	2, 2, 194, 195, 3, 2, 2, 2, 195, 197, 5, 12, 7, 2, 196, 193, 3, 2, 2, 2,
	197, 198, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199,
	201, 3, 2, 2, 2, 200, 202, 5, 6, 4, 2, 201, 200, 3, 2, 2, 2, 201, 202,
	3, 2, 2, 2, 202, 205, 3, 2, 2, 2, 203, 205, 5, 6, 4, 2, 204, 196, 3, 2,
	2, 2, 204, 203, 3, 2, 2, 2, 205, 15, 3, 2, 2, 2, 206, 207, 9, 2, 2, 2,
	207, 17, 3, 2, 2, 2, 208, 210, 5, 16, 9, 2, 209, 208, 3, 2, 2, 2, 210,
	211, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 19, 3,
	2, 2, 2, 213, 215, 5, 16, 9, 2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2,
	2, 216, 214, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 226, 3, 2, 2, 2, 218,
	220, 7, 25, 2, 2, 219, 221, 5, 16, 9, 2, 220, 219, 3, 2, 2, 2, 221, 222,
	3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 225, 3, 2,
	2, 2, 224, 218, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2,
	226, 227, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229,
	231, 7, 25, 2, 2, 230, 229, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 21,
	3, 2, 2, 2, 232, 268, 7, 12, 2, 2, 233, 268, 7, 14, 2, 2, 234, 268, 7,
	15, 2, 2, 235, 268, 7, 16, 2, 2, 236, 268, 7, 17, 2, 2, 237, 268, 7, 18,
	2, 2, 238, 268, 7, 19, 2, 2, 239, 268, 7, 20, 2, 2, 240, 268, 7, 21, 2,
	2, 241, 268, 7, 22, 2, 2, 242, 268, 7, 23, 2, 2, 243, 268, 7, 24, 2, 2,
	244, 268, 7, 25, 2, 2, 245, 268, 7, 26, 2, 2, 246, 268, 7, 27, 2, 2, 247,
	268, 7, 28, 2, 2, 248, 268, 7, 29, 2, 2, 249, 268, 7, 30, 2, 2, 250, 268,
	7, 31, 2, 2, 251, 268, 7, 32, 2, 2, 252, 268, 7, 33, 2, 2, 253, 268, 7,
	34, 2, 2, 254, 268, 7, 35, 2, 2, 255, 268, 7, 36, 2, 2, 256, 268, 7, 38,
	2, 2, 257, 268, 7, 39, 2, 2, 258, 268, 7, 40, 2, 2, 259, 268, 7, 41, 2,
	2, 260, 268, 7, 42, 2, 2, 261, 268, 7, 43, 2, 2, 262, 268, 7, 44, 2, 2,
	263, 268, 7, 45, 2, 2, 264, 268, 7, 46, 2, 2, 265, 268, 5, 68, 35, 2, 266,
	268, 7, 48, 2, 2, 267, 232, 3, 2, 2, 2, 267, 233, 3, 2, 2, 2, 267, 234,
	3, 2, 2, 2, 267, 235, 3, 2, 2, 2, 267, 236, 3, 2, 2, 2, 267, 237, 3, 2,
	2, 2, 267, 238, 3, 2, 2, 2, 267, 239, 3, 2, 2, 2, 267, 240, 3, 2, 2, 2,
	267, 241, 3, 2, 2, 2, 267, 242, 3, 2, 2, 2, 267, 243, 3, 2, 2, 2, 267,
	244, 3, 2, 2, 2, 267, 245, 3, 2, 2, 2, 267, 246, 3, 2, 2, 2, 267, 247,
	3, 2, 2, 2, 267, 248, 3, 2, 2, 2, 267, 249, 3, 2, 2, 2, 267, 250, 3, 2,
	2, 2, 267, 251, 3, 2, 2, 2, 267, 252, 3, 2, 2, 2, 267, 253, 3, 2, 2, 2,
	267, 254, 3, 2, 2, 2, 267, 255, 3, 2, 2, 2, 267, 256, 3, 2, 2, 2, 267,
	257, 3, 2, 2, 2, 267, 258, 3, 2, 2, 2, 267, 259, 3, 2, 2, 2, 267, 260,
	3, 2, 2, 2, 267, 261, 3, 2, 2, 2, 267, 262, 3, 2, 2, 2, 267, 263, 3, 2,
	2, 2, 267, 264, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 267, 266, 3, 2, 2, 2,
	268, 23, 3, 2, 2, 2, 269, 272, 5, 22, 12, 2, 270, 272, 5, 4, 3, 2, 271,
	269, 3, 2, 2, 2, 271, 270, 3, 2, 2, 2, 272, 25, 3, 2, 2, 2, 273, 275, 5,
	6, 4, 2, 274, 273, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 276, 3, 2, 2,
	2, 276, 278, 5, 24, 13, 2, 277, 274, 3, 2, 2, 2, 278, 281, 3, 2, 2, 2,
	279, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 27, 3, 2, 2, 2, 281, 279,
	3, 2, 2, 2, 282, 283, 7, 13, 2, 2, 283, 285, 5, 26, 14, 2, 284, 286, 5,
	6, 4, 2, 285, 284, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 3, 2, 2,
	2, 287, 288, 7, 13, 2, 2, 288, 29, 3, 2, 2, 2, 289, 291, 5, 14, 8, 2, 290,
	289, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 294,
	5, 90, 46, 2, 293, 295, 5, 14, 8, 2, 294, 293, 3, 2, 2, 2, 294, 295, 3,
	2, 2, 2, 295, 312, 3, 2, 2, 2, 296, 298, 5, 14, 8, 2, 297, 296, 3, 2, 2,
	2, 297, 298, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 301, 5, 18, 10, 2,
	300, 302, 5, 14, 8, 2, 301, 300, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302,
	312, 3, 2, 2, 2, 303, 305, 5, 14, 8, 2, 304, 303, 3, 2, 2, 2, 304, 305,
	3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 308, 5, 28, 15, 2, 307, 309, 5,
	14, 8, 2, 308, 307, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 312, 3, 2, 2,
	2, 310, 312, 7, 34, 2, 2, 311, 290, 3, 2, 2, 2, 311, 297, 3, 2, 2, 2, 311,
	304, 3, 2, 2, 2, 311, 310, 3, 2, 2, 2, 312, 31, 3, 2, 2, 2, 313, 316, 5,
	34, 18, 2, 314, 316, 5, 40, 21, 2, 315, 313, 3, 2, 2, 2, 315, 314, 3, 2,
	2, 2, 316, 33, 3, 2, 2, 2, 317, 320, 5, 36, 19, 2, 318, 320, 5, 50, 26,
	2, 319, 317, 3, 2, 2, 2, 319, 318, 3, 2, 2, 2, 320, 35, 3, 2, 2, 2, 321,
	323, 5, 42, 22, 2, 322, 321, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 324,
	3, 2, 2, 2, 324, 325, 5, 38, 20, 2, 325, 37, 3, 2, 2, 2, 326, 328, 5, 14,
	8, 2, 327, 326, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2,
	329, 331, 7, 30, 2, 2, 330, 332, 5, 50, 26, 2, 331, 330, 3, 2, 2, 2, 331,
	332, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 335, 7, 32, 2, 2, 334, 336,
	5, 14, 8, 2, 335, 334, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 339, 3, 2,
	2, 2, 337, 339, 5, 74, 38, 2, 338, 327, 3, 2, 2, 2, 338, 337, 3, 2, 2,
	2, 339, 39, 3, 2, 2, 2, 340, 341, 5, 42, 22, 2, 341, 343, 7, 28, 2, 2,
	342, 344, 5, 48, 25, 2, 343, 342, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344,
	345, 3, 2, 2, 2, 345, 347, 7, 29, 2, 2, 346, 348, 5, 14, 8, 2, 347, 346,
	3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 41, 3, 2, 2, 2, 349, 351, 5, 30,
	16, 2, 350, 349, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2,
	352, 353, 3, 2, 2, 2, 353, 356, 3, 2, 2, 2, 354, 356, 5, 64, 33, 2, 355,
	350, 3, 2, 2, 2, 355, 354, 3, 2, 2, 2, 356, 43, 3, 2, 2, 2, 357, 362, 5,
	34, 18, 2, 358, 359, 7, 23, 2, 2, 359, 361, 5, 34, 18, 2, 360, 358, 3,
	2, 2, 2, 361, 364, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2,
	2, 363, 367, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 365, 367, 5, 80, 41, 2,
	366, 357, 3, 2, 2, 2, 366, 365, 3, 2, 2, 2, 367, 45, 3, 2, 2, 2, 368, 373,
	5, 32, 17, 2, 369, 370, 7, 23, 2, 2, 370, 372, 5, 32, 17, 2, 371, 369,
	3, 2, 2, 2, 372, 375, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 373, 374, 3, 2,
	2, 2, 374, 376, 3, 2, 2, 2, 375, 373, 3, 2, 2, 2, 376, 377, 7, 2, 2, 3,
	377, 382, 3, 2, 2, 2, 378, 379, 5, 82, 42, 2, 379, 380, 7, 2, 2, 3, 380,
	382, 3, 2, 2, 2, 381, 368, 3, 2, 2, 2, 381, 378, 3, 2, 2, 2, 382, 47, 3,
	2, 2, 2, 383, 387, 5, 44, 23, 2, 384, 387, 5, 14, 8, 2, 385, 387, 5, 84,
	43, 2, 386, 383, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 385, 3, 2, 2, 2,
	387, 49, 3, 2, 2, 2, 388, 389, 5, 52, 27, 2, 389, 390, 7, 34, 2, 2, 390,
	393, 5, 56, 29, 2, 391, 392, 7, 28, 2, 2, 392, 394, 5, 54, 28, 2, 393,
	391, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 51, 3, 2, 2, 2, 395, 397, 5,
	14, 8, 2, 396, 395, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 398, 3, 2, 2,
	2, 398, 400, 5, 20, 11, 2, 399, 401, 5, 14, 8, 2, 400, 399, 3, 2, 2, 2,
	400, 401, 3, 2, 2, 2, 401, 411, 3, 2, 2, 2, 402, 404, 5, 14, 8, 2, 403,
	402, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 407,
	5, 28, 15, 2, 406, 408, 5, 14, 8, 2, 407, 406, 3, 2, 2, 2, 407, 408, 3,
	2, 2, 2, 408, 411, 3, 2, 2, 2, 409, 411, 5, 86, 44, 2, 410, 396, 3, 2,
	2, 2, 410, 403, 3, 2, 2, 2, 410, 409, 3, 2, 2, 2, 411, 53, 3, 2, 2, 2,
	412, 414, 7, 27, 2, 2, 413, 412, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415,
	413, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 55, 3, 2, 2, 2, 417, 419, 5,
	14, 8, 2, 418, 417, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420, 3, 2, 2,
	2, 420, 422, 5, 20, 11, 2, 421, 423, 5, 14, 8, 2, 422, 421, 3, 2, 2, 2,
	422, 423, 3, 2, 2, 2, 423, 439, 3, 2, 2, 2, 424, 426, 5, 14, 8, 2, 425,
	424, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 427, 3, 2, 2, 2, 427, 429,
	5, 58, 30, 2, 428, 430, 5, 14, 8, 2, 429, 428, 3, 2, 2, 2, 429, 430, 3,
	2, 2, 2, 430, 439, 3, 2, 2, 2, 431, 433, 5, 14, 8, 2, 432, 431, 3, 2, 2,
	2, 432, 433, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 436, 5, 88, 45, 2,
	435, 437, 5, 14, 8, 2, 436, 435, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437,
	439, 3, 2, 2, 2, 438, 418, 3, 2, 2, 2, 438, 425, 3, 2, 2, 2, 438, 432,
	3, 2, 2, 2, 439, 57, 3, 2, 2, 2, 440, 447, 7, 36, 2, 2, 441, 443, 5, 6,
	4, 2, 442, 441, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2,
	444, 446, 5, 60, 31, 2, 445, 442, 3, 2, 2, 2, 446, 449, 3, 2, 2, 2, 447,
	445, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 451, 3, 2, 2, 2, 449, 447,
	3, 2, 2, 2, 450, 452, 5, 6, 4, 2, 451, 450, 3, 2, 2, 2, 451, 452, 3, 2,
	2, 2, 452, 453, 3, 2, 2, 2, 453, 454, 7, 38, 2, 2, 454, 59, 3, 2, 2, 2,
	455, 456, 9, 3, 2, 2, 456, 61, 3, 2, 2, 2, 457, 458, 9, 4, 2, 2, 458, 63,
	3, 2, 2, 2, 459, 465, 5, 30, 16, 2, 460, 464, 5, 30, 16, 2, 461, 464, 7,
	25, 2, 2, 462, 464, 5, 14, 8, 2, 463, 460, 3, 2, 2, 2, 463, 461, 3, 2,
	2, 2, 463, 462, 3, 2, 2, 2, 464, 467, 3, 2, 2, 2, 465, 463, 3, 2, 2, 2,
	465, 466, 3, 2, 2, 2, 466, 65, 3, 2, 2, 2, 467, 465, 3, 2, 2, 2, 468, 469,
	5, 62, 32, 2, 469, 67, 3, 2, 2, 2, 470, 471, 5, 62, 32, 2, 471, 69, 3,
	2, 2, 2, 472, 477, 7, 37, 2, 2, 473, 478, 7, 3, 2, 2, 474, 478, 5, 62,
	32, 2, 475, 478, 7, 6, 2, 2, 476, 478, 7, 9, 2, 2, 477, 473, 3, 2, 2, 2,
	477, 474, 3, 2, 2, 2, 477, 475, 3, 2, 2, 2, 477, 476, 3, 2, 2, 2, 478,
	71, 3, 2, 2, 2, 479, 481, 5, 106, 54, 2, 480, 479, 3, 2, 2, 2, 481, 482,
	3, 2, 2, 2, 482, 480, 3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 484, 3, 2,
	2, 2, 484, 486, 5, 104, 53, 2, 485, 487, 5, 106, 54, 2, 486, 485, 3, 2,
	2, 2, 487, 488, 3, 2, 2, 2, 488, 486, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2,
	489, 73, 3, 2, 2, 2, 490, 492, 5, 14, 8, 2, 491, 490, 3, 2, 2, 2, 491,
	492, 3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493, 494, 7, 30, 2, 2, 494, 495,
	5, 76, 39, 2, 495, 496, 5, 50, 26, 2, 496, 498, 7, 32, 2, 2, 497, 499,
	5, 14, 8, 2, 498, 497, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 75, 3, 2,
	2, 2, 500, 501, 5, 78, 40, 2, 501, 502, 7, 28, 2, 2, 502, 77, 3, 2, 2,
	2, 503, 506, 5, 14, 8, 2, 504, 506, 7, 23, 2, 2, 505, 503, 3, 2, 2, 2,
	505, 504, 3, 2, 2, 2, 506, 509, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2, 507,
	508, 3, 2, 2, 2, 508, 510, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 510, 511,
	7, 34, 2, 2, 511, 522, 5, 56, 29, 2, 512, 514, 7, 23, 2, 2, 513, 515, 5,
	14, 8, 2, 514, 513, 3, 2, 2, 2, 514, 515, 3, 2, 2, 2, 515, 518, 3, 2, 2,
	2, 516, 517, 7, 34, 2, 2, 517, 519, 5, 56, 29, 2, 518, 516, 3, 2, 2, 2,
	518, 519, 3, 2, 2, 2, 519, 521, 3, 2, 2, 2, 520, 512, 3, 2, 2, 2, 521,
	524, 3, 2, 2, 2, 522, 520, 3, 2, 2, 2, 522, 523, 3, 2, 2, 2, 523, 79, 3,
	2, 2, 2, 524, 522, 3, 2, 2, 2, 525, 527, 5, 14, 8, 2, 526, 525, 3, 2, 2,
	2, 526, 527, 3, 2, 2, 2, 527, 528, 3, 2, 2, 2, 528, 530, 7, 23, 2, 2, 529,
	526, 3, 2, 2, 2, 530, 533, 3, 2, 2, 2, 531, 529, 3, 2, 2, 2, 531, 532,
	3, 2, 2, 2, 532, 534, 3, 2, 2, 2, 533, 531, 3, 2, 2, 2, 534, 542, 5, 34,
	18, 2, 535, 538, 7, 23, 2, 2, 536, 539, 5, 34, 18, 2, 537, 539, 5, 14,
	8, 2, 538, 536, 3, 2, 2, 2, 538, 537, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2,
	539, 541, 3, 2, 2, 2, 540, 535, 3, 2, 2, 2, 541, 544, 3, 2, 2, 2, 542,
	540, 3, 2, 2, 2, 542, 543, 3, 2, 2, 2, 543, 81, 3, 2, 2, 2, 544, 542, 3,
	2, 2, 2, 545, 547, 5, 14, 8, 2, 546, 545, 3, 2, 2, 2, 546, 547, 3, 2, 2,
	2, 547, 548, 3, 2, 2, 2, 548, 550, 7, 23, 2, 2, 549, 546, 3, 2, 2, 2, 550,
	553, 3, 2, 2, 2, 551, 549, 3, 2, 2, 2, 551, 552, 3, 2, 2, 2, 552, 554,
	3, 2, 2, 2, 553, 551, 3, 2, 2, 2, 554, 562, 5, 32, 17, 2, 555, 558, 7,
	23, 2, 2, 556, 559, 5, 32, 17, 2, 557, 559, 5, 14, 8, 2, 558, 556, 3, 2,
	2, 2, 558, 557, 3, 2, 2, 2, 558, 559, 3, 2, 2, 2, 559, 561, 3, 2, 2, 2,
	560, 555, 3, 2, 2, 2, 561, 564, 3, 2, 2, 2, 562, 560, 3, 2, 2, 2, 562,
	563, 3, 2, 2, 2, 563, 83, 3, 2, 2, 2, 564, 562, 3, 2, 2, 2, 565, 567, 5,
	14, 8, 2, 566, 565, 3, 2, 2, 2, 566, 567, 3, 2, 2, 2, 567, 568, 3, 2, 2,
	2, 568, 570, 7, 23, 2, 2, 569, 566, 3, 2, 2, 2, 570, 571, 3, 2, 2, 2, 571,
	569, 3, 2, 2, 2, 571, 572, 3, 2, 2, 2, 572, 574, 3, 2, 2, 2, 573, 575,
	5, 14, 8, 2, 574, 573, 3, 2, 2, 2, 574, 575, 3, 2, 2, 2, 575, 85, 3, 2,
	2, 2, 576, 581, 5, 30, 16, 2, 577, 578, 7, 25, 2, 2, 578, 580, 5, 30, 16,
	2, 579, 577, 3, 2, 2, 2, 580, 583, 3, 2, 2, 2, 581, 579, 3, 2, 2, 2, 581,
	582, 3, 2, 2, 2, 582, 87, 3, 2, 2, 2, 583, 581, 3, 2, 2, 2, 584, 589, 5,
	18, 10, 2, 585, 586, 7, 25, 2, 2, 586, 588, 5, 18, 10, 2, 587, 585, 3,
	2, 2, 2, 588, 591, 3, 2, 2, 2, 589, 587, 3, 2, 2, 2, 589, 590, 3, 2, 2,
	2, 590, 89, 3, 2, 2, 2, 591, 589, 3, 2, 2, 2, 592, 593, 7, 31, 2, 2, 593,
	594, 7, 33, 2, 2, 594, 595, 5, 92, 47, 2, 595, 596, 7, 33, 2, 2, 596, 597,
	5, 94, 48, 2, 597, 598, 7, 33, 2, 2, 598, 599, 5, 100, 51, 2, 599, 600,
	7, 33, 2, 2, 600, 601, 7, 31, 2, 2, 601, 91, 3, 2, 2, 2, 602, 603, 5, 96,
	49, 2, 603, 93, 3, 2, 2, 2, 604, 605, 5, 96, 49, 2, 605, 95, 3, 2, 2, 2,
	606, 608, 5, 98, 50, 2, 607, 606, 3, 2, 2, 2, 608, 609, 3, 2, 2, 2, 609,
	607, 3, 2, 2, 2, 609, 610, 3, 2, 2, 2, 610, 97, 3, 2, 2, 2, 611, 612, 9,
	5, 2, 2, 612, 99, 3, 2, 2, 2, 613, 615, 5, 102, 52, 2, 614, 613, 3, 2,
	2, 2, 615, 616, 3, 2, 2, 2, 616, 614, 3, 2, 2, 2, 616, 617, 3, 2, 2, 2,
	617, 101, 3, 2, 2, 2, 618, 619, 9, 6, 2, 2, 619, 103, 3, 2, 2, 2, 620,
	621, 7, 9, 2, 2, 621, 622, 7, 6, 2, 2, 622, 105, 3, 2, 2, 2, 623, 624,
	9, 7, 2, 2, 624, 107, 3, 2, 2, 2, 625, 626, 9, 8, 2, 2, 626, 109, 3, 2,
	2, 2, 94, 112, 117, 122, 126, 131, 134, 170, 175, 179, 184, 188, 193, 198,
	201, 204, 211, 216, 222, 226, 230, 267, 271, 274, 279, 285, 290, 294, 297,
	301, 304, 308, 311, 315, 319, 322, 327, 331, 335, 338, 343, 347, 352, 355,
	362, 366, 373, 381, 386, 393, 396, 400, 403, 407, 410, 415, 418, 422, 425,
	429, 432, 436, 438, 442, 447, 451, 463, 465, 477, 482, 488, 491, 498, 505,
	507, 514, 518, 522, 526, 531, 538, 542, 546, 551, 558, 562, 566, 571, 574,
	581, 589, 609, 616,
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
	"mailboxList", "addressList", "groupList", "addrSpec", "localPart", "port",
	"domain", "domainLiteral", "dtext", "obsNoWSCTL", "obsPhrase", "obsCtext",
	"obsQtext", "obsQP", "obsFWS", "obsAngleAddr", "obsRoute", "obsDomainList",
	"obsMboxList", "obsAddrList", "obsGroupList", "obsLocalPart", "obsDomain",
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
	AddressParserRULE_port          = 26
	AddressParserRULE_domain        = 27
	AddressParserRULE_domainLiteral = 28
	AddressParserRULE_dtext         = 29
	AddressParserRULE_obsNoWSCTL    = 30
	AddressParserRULE_obsPhrase     = 31
	AddressParserRULE_obsCtext      = 32
	AddressParserRULE_obsQtext      = 33
	AddressParserRULE_obsQP         = 34
	AddressParserRULE_obsFWS        = 35
	AddressParserRULE_obsAngleAddr  = 36
	AddressParserRULE_obsRoute      = 37
	AddressParserRULE_obsDomainList = 38
	AddressParserRULE_obsMboxList   = 39
	AddressParserRULE_obsAddrList   = 40
	AddressParserRULE_obsGroupList  = 41
	AddressParserRULE_obsLocalPart  = 42
	AddressParserRULE_obsDomain     = 43
	AddressParserRULE_encodedWord   = 44
	AddressParserRULE_charset       = 45
	AddressParserRULE_encoding      = 46
	AddressParserRULE_token         = 47
	AddressParserRULE_tokenChar     = 48
	AddressParserRULE_encodedText   = 49
	AddressParserRULE_encodedChar   = 50
	AddressParserRULE_crlf          = 51
	AddressParserRULE_wsp           = 52
	AddressParserRULE_vchar         = 53
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

	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation, AddressParserDQuote, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserLParens, AddressParserRParens, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserBackslash, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde, AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Vchar()
		}

	case AddressParserTAB, AddressParserSP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(AddressParserBackslash)
		}
		{
			p.SetState(113)
			p.QuotedChar()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(124)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AddressParserTAB || _la == AddressParserSP {
				{
					p.SetState(117)
					p.Wsp()
				}

				p.SetState(122)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(123)
				p.Crlf()
			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(126)
					p.Wsp()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
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

	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(AddressParserExclamation)
		}

	case AddressParserDQuote:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(AddressParserDQuote)
		}

	case AddressParserHash:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.Match(AddressParserHash)
		}

	case AddressParserDollar:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.Match(AddressParserDollar)
		}

	case AddressParserPercent:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			p.Match(AddressParserPercent)
		}

	case AddressParserAmpersand:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(139)
			p.Match(AddressParserAmpersand)
		}

	case AddressParserSQuote:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(140)
			p.Match(AddressParserSQuote)
		}

	case AddressParserAsterisk:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(141)
			p.Match(AddressParserAsterisk)
		}

	case AddressParserPlus:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(142)
			p.Match(AddressParserPlus)
		}

	case AddressParserComma:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(143)
			p.Match(AddressParserComma)
		}

	case AddressParserMinus:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(144)
			p.Match(AddressParserMinus)
		}

	case AddressParserPeriod:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(145)
			p.Match(AddressParserPeriod)
		}

	case AddressParserSlash:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(146)
			p.Match(AddressParserSlash)
		}

	case AddressParserDigit:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(147)
			p.Match(AddressParserDigit)
		}

	case AddressParserColon:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(148)
			p.Match(AddressParserColon)
		}

	case AddressParserSemicolon:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(149)
			p.Match(AddressParserSemicolon)
		}

	case AddressParserLess:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(150)
			p.Match(AddressParserLess)
		}

	case AddressParserEqual:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(151)
			p.Match(AddressParserEqual)
		}

	case AddressParserGreater:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(152)
			p.Match(AddressParserGreater)
		}

	case AddressParserQuestion:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(153)
			p.Match(AddressParserQuestion)
		}

	case AddressParserAt:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(154)
			p.Match(AddressParserAt)
		}

	case AddressParserAlphaUpper:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(155)
			p.Match(AddressParserAlphaUpper)
		}

	case AddressParserLBracket:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(156)
			p.Match(AddressParserLBracket)
		}

	case AddressParserRBracket:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(157)
			p.Match(AddressParserRBracket)
		}

	case AddressParserCaret:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(158)
			p.Match(AddressParserCaret)
		}

	case AddressParserUnderscore:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(159)
			p.Match(AddressParserUnderscore)
		}

	case AddressParserBacktick:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(160)
			p.Match(AddressParserBacktick)
		}

	case AddressParserAlphaLower:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(161)
			p.Match(AddressParserAlphaLower)
		}

	case AddressParserLCurly:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(162)
			p.Match(AddressParserLCurly)
		}

	case AddressParserPipe:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(163)
			p.Match(AddressParserPipe)
		}

	case AddressParserRCurly:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(164)
			p.Match(AddressParserRCurly)
		}

	case AddressParserTilde:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(165)
			p.Match(AddressParserTilde)
		}

	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserDelete:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(166)
			p.ObsCtext()
		}

	case AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(167)
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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserExclamation, AddressParserDQuote, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde, AddressParserDelete, AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Ctext()
		}

	case AddressParserBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.QuotedPair()
		}

	case AddressParserLParens:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(172)
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
		p.SetState(175)
		p.Match(AddressParserLParens)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
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
				p.Ccontent()
			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(185)
			p.Fws()
		}

	}
	{
		p.SetState(188)
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

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(191)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
					{
						p.SetState(190)
						p.Fws()
					}

				}
				{
					p.SetState(193)
					p.Comment()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(198)
				p.Fws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
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
		p.SetState(204)
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
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(206)
				p.Atext()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(209)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33))|(1<<(AddressParserUTF8NonAscii-33)))) != 0) {
		{
			p.SetState(211)
			p.Atext()
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(216)
				p.Match(AddressParserPeriod)
			}
			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33))|(1<<(AddressParserUTF8NonAscii-33)))) != 0) {
				{
					p.SetState(217)
					p.Atext()
				}

				p.SetState(220)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AddressParserPeriod {
		{
			p.SetState(227)
			p.Match(AddressParserPeriod)
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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserExclamation:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Match(AddressParserExclamation)
		}

	case AddressParserHash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.Match(AddressParserHash)
		}

	case AddressParserDollar:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(232)
			p.Match(AddressParserDollar)
		}

	case AddressParserPercent:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(233)
			p.Match(AddressParserPercent)
		}

	case AddressParserAmpersand:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(234)
			p.Match(AddressParserAmpersand)
		}

	case AddressParserSQuote:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(235)
			p.Match(AddressParserSQuote)
		}

	case AddressParserLParens:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(236)
			p.Match(AddressParserLParens)
		}

	case AddressParserRParens:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(237)
			p.Match(AddressParserRParens)
		}

	case AddressParserAsterisk:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(238)
			p.Match(AddressParserAsterisk)
		}

	case AddressParserPlus:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(239)
			p.Match(AddressParserPlus)
		}

	case AddressParserComma:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(240)
			p.Match(AddressParserComma)
		}

	case AddressParserMinus:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(241)
			p.Match(AddressParserMinus)
		}

	case AddressParserPeriod:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(242)
			p.Match(AddressParserPeriod)
		}

	case AddressParserSlash:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(243)
			p.Match(AddressParserSlash)
		}

	case AddressParserDigit:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(244)
			p.Match(AddressParserDigit)
		}

	case AddressParserColon:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(245)
			p.Match(AddressParserColon)
		}

	case AddressParserSemicolon:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(246)
			p.Match(AddressParserSemicolon)
		}

	case AddressParserLess:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(247)
			p.Match(AddressParserLess)
		}

	case AddressParserEqual:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(248)
			p.Match(AddressParserEqual)
		}

	case AddressParserGreater:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(249)
			p.Match(AddressParserGreater)
		}

	case AddressParserQuestion:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(250)
			p.Match(AddressParserQuestion)
		}

	case AddressParserAt:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(251)
			p.Match(AddressParserAt)
		}

	case AddressParserAlphaUpper:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(252)
			p.Match(AddressParserAlphaUpper)
		}

	case AddressParserLBracket:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(253)
			p.Match(AddressParserLBracket)
		}

	case AddressParserRBracket:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(254)
			p.Match(AddressParserRBracket)
		}

	case AddressParserCaret:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(255)
			p.Match(AddressParserCaret)
		}

	case AddressParserUnderscore:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(256)
			p.Match(AddressParserUnderscore)
		}

	case AddressParserBacktick:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(257)
			p.Match(AddressParserBacktick)
		}

	case AddressParserAlphaLower:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(258)
			p.Match(AddressParserAlphaLower)
		}

	case AddressParserLCurly:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(259)
			p.Match(AddressParserLCurly)
		}

	case AddressParserPipe:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(260)
			p.Match(AddressParserPipe)
		}

	case AddressParserRCurly:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(261)
			p.Match(AddressParserRCurly)
		}

	case AddressParserTilde:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(262)
			p.Match(AddressParserTilde)
		}

	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserDelete:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(263)
			p.ObsQtext()
		}

	case AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(264)
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

	p.SetState(269)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserExclamation, AddressParserHash, AddressParserDollar, AddressParserPercent, AddressParserAmpersand, AddressParserSQuote, AddressParserLParens, AddressParserRParens, AddressParserAsterisk, AddressParserPlus, AddressParserComma, AddressParserMinus, AddressParserPeriod, AddressParserSlash, AddressParserDigit, AddressParserColon, AddressParserSemicolon, AddressParserLess, AddressParserEqual, AddressParserGreater, AddressParserQuestion, AddressParserAt, AddressParserAlphaUpper, AddressParserLBracket, AddressParserRBracket, AddressParserCaret, AddressParserUnderscore, AddressParserBacktick, AddressParserAlphaLower, AddressParserLCurly, AddressParserPipe, AddressParserRCurly, AddressParserTilde, AddressParserDelete, AddressParserUTF8NonAscii:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.Qtext()
		}

	case AddressParserBackslash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(268)
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
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(272)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(271)
					p.Fws()
				}

			}
			{
				p.SetState(274)
				p.QuotedContent()
			}

		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
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
		p.SetState(280)
		p.Match(AddressParserDQuote)
	}
	{
		p.SetState(281)
		p.QuotedValue()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(282)
			p.Fws()
		}

	}
	{
		p.SetState(285)
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

func (s *WordContext) At() antlr.TerminalNode {
	return s.GetToken(AddressParserAt, 0)
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

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(287)
				p.Cfws()
			}

		}
		{
			p.SetState(290)
			p.EncodedWord()
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(291)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(294)
				p.Cfws()
			}

		}
		{
			p.SetState(297)
			p.Atom()
		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(298)
				p.Cfws()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(301)
				p.Cfws()
			}

		}
		{
			p.SetState(304)
			p.QuotedString()
		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(305)
				p.Cfws()
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(308)
			p.Match(AddressParserAt)
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

	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)
			p.Mailbox()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(312)
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

	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.NameAddr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(316)
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
	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(319)
			p.DisplayName()
		}

	}
	{
		p.SetState(322)
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

	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(324)
				p.Cfws()
			}

		}
		{
			p.SetState(327)
			p.Match(AddressParserLess)
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserAt-32))|(1<<(AddressParserAlphaUpper-32))|(1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32))|(1<<(AddressParserUTF8NonAscii-32)))) != 0) {
			{
				p.SetState(328)
				p.AddrSpec()
			}

		}
		{
			p.SetState(331)
			p.Match(AddressParserGreater)
		}
		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(332)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(335)
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
		p.SetState(338)
		p.DisplayName()
	}
	{
		p.SetState(339)
		p.Match(AddressParserColon)
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserQuestion))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserAt-32))|(1<<(AddressParserAlphaUpper-32))|(1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32))|(1<<(AddressParserUTF8NonAscii-32)))) != 0) {
		{
			p.SetState(340)
			p.GroupList()
		}

	}
	{
		p.SetState(343)
		p.Match(AddressParserSemicolon)
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(344)
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

	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(347)
					p.Word()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(350)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
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

	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.Mailbox()
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AddressParserComma {
			{
				p.SetState(356)
				p.Match(AddressParserComma)
			}
			{
				p.SetState(357)
				p.Mailbox()
			}

			p.SetState(362)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
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

func (s *AddressListContext) EOF() antlr.TerminalNode {
	return s.GetToken(AddressParserEOF, 0)
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

	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)
			p.Address()
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AddressParserComma {
			{
				p.SetState(367)
				p.Match(AddressParserComma)
			}
			{
				p.SetState(368)
				p.Address()
			}

			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(374)
			p.Match(AddressParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(376)
			p.ObsAddrList()
		}
		{
			p.SetState(377)
			p.Match(AddressParserEOF)
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

	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(381)
			p.MailboxList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(382)
			p.Cfws()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(383)
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

func (s *AddrSpecContext) Colon() antlr.TerminalNode {
	return s.GetToken(AddressParserColon, 0)
}

func (s *AddrSpecContext) Port() IPortContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortContext)
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
		p.SetState(386)
		p.LocalPart()
	}
	{
		p.SetState(387)
		p.Match(AddressParserAt)
	}
	{
		p.SetState(388)
		p.Domain()
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AddressParserColon {
		{
			p.SetState(389)
			p.Match(AddressParserColon)
		}
		{
			p.SetState(390)
			p.Port()
		}

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

	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(393)
				p.Cfws()
			}

		}
		{
			p.SetState(396)
			p.DotAtom()
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(397)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
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
			p.QuotedString()
		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(404)
				p.Cfws()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(407)
			p.ObsLocalPart()
		}

	}

	return localctx
}

// IPortContext is an interface to support dynamic dispatch.
type IPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortContext differentiates from other interfaces.
	IsPortContext()
}

type PortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortContext() *PortContext {
	var p = new(PortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AddressParserRULE_port
	return p
}

func (*PortContext) IsPortContext() {}

func NewPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortContext {
	var p = new(PortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AddressParserRULE_port

	return p
}

func (s *PortContext) GetParser() antlr.Parser { return s.parser }

func (s *PortContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(AddressParserDigit)
}

func (s *PortContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(AddressParserDigit, i)
}

func (s *PortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.EnterPort(s)
	}
}

func (s *PortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AddressParserListener); ok {
		listenerT.ExitPort(s)
	}
}

func (p *AddressParser) Port() (localctx IPortContext) {
	localctx = NewPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AddressParserRULE_port)
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
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AddressParserDigit {
		{
			p.SetState(410)
			p.Match(AddressParserDigit)
		}

		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(415)
				p.Cfws()
			}

		}
		{
			p.SetState(418)
			p.DotAtom()
		}
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(419)
				p.Cfws()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(422)
				p.Cfws()
			}

		}
		{
			p.SetState(425)
			p.DomainLiteral()
		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(426)
				p.Cfws()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(429)
				p.Cfws()
			}

		}
		{
			p.SetState(432)
			p.ObsDomain()
		}
		p.SetState(434)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(433)
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
		p.SetState(438)
		p.Match(AddressParserLBracket)
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(440)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
				{
					p.SetState(439)
					p.Fws()
				}

			}
			{
				p.SetState(442)
				p.Dtext()
			}

		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())
	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP))) != 0 {
		{
			p.SetState(448)
			p.Fws()
		}

	}
	{
		p.SetState(451)
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
		p.SetState(453)
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
	p.EnterRule(localctx, 60, AddressParserRULE_obsNoWSCTL)
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
		p.SetState(455)
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
	p.EnterRule(localctx, 62, AddressParserRULE_obsPhrase)

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
		p.SetState(457)
		p.Word()
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(461)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(458)
					p.Word()
				}

			case 2:
				{
					p.SetState(459)
					p.Match(AddressParserPeriod)
				}

			case 3:
				{
					p.SetState(460)
					p.Cfws()
				}

			}

		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 64, AddressParserRULE_obsCtext)

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
		p.SetState(466)
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
	p.EnterRule(localctx, 66, AddressParserRULE_obsQtext)

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
		p.SetState(468)
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
	p.EnterRule(localctx, 68, AddressParserRULE_obsQP)

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
		p.Match(AddressParserBackslash)
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AddressParserU_00:
		{
			p.SetState(471)
			p.Match(AddressParserU_00)
		}

	case AddressParserU_01_08, AddressParserU_0B, AddressParserU_0C, AddressParserU_0E_1F, AddressParserDelete:
		{
			p.SetState(472)
			p.ObsNoWSCTL()
		}

	case AddressParserLF:
		{
			p.SetState(473)
			p.Match(AddressParserLF)
		}

	case AddressParserCR:
		{
			p.SetState(474)
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
	p.EnterRule(localctx, 70, AddressParserRULE_obsFWS)
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
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AddressParserTAB || _la == AddressParserSP {
		{
			p.SetState(477)
			p.Wsp()
		}

		p.SetState(480)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(482)
		p.Crlf()
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(483)
				p.Wsp()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 72, AddressParserRULE_obsAngleAddr)
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
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(488)
			p.Cfws()
		}

	}
	{
		p.SetState(491)
		p.Match(AddressParserLess)
	}
	{
		p.SetState(492)
		p.ObsRoute()
	}
	{
		p.SetState(493)
		p.AddrSpec()
	}
	{
		p.SetState(494)
		p.Match(AddressParserGreater)
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(495)
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
	p.EnterRule(localctx, 74, AddressParserRULE_obsRoute)

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
		p.SetState(498)
		p.ObsDomainList()
	}
	{
		p.SetState(499)
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
	p.EnterRule(localctx, 76, AddressParserRULE_obsDomainList)
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
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens)|(1<<AddressParserComma))) != 0 {
		p.SetState(503)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AddressParserTAB, AddressParserCR, AddressParserSP, AddressParserLParens:
			{
				p.SetState(501)
				p.Cfws()
			}

		case AddressParserComma:
			{
				p.SetState(502)
				p.Match(AddressParserComma)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(508)
		p.Match(AddressParserAt)
	}
	{
		p.SetState(509)
		p.Domain()
	}
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(510)
			p.Match(AddressParserComma)
		}
		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
			{
				p.SetState(511)
				p.Cfws()
			}

		}
		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AddressParserAt {
			{
				p.SetState(514)
				p.Match(AddressParserAt)
			}
			{
				p.SetState(515)
				p.Domain()
			}

		}

		p.SetState(522)
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
	p.EnterRule(localctx, 78, AddressParserRULE_obsMboxList)
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
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(524)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(523)
					p.Cfws()
				}

			}
			{
				p.SetState(526)
				p.Match(AddressParserComma)
			}

		}
		p.SetState(531)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext())
	}
	{
		p.SetState(532)
		p.Mailbox()
	}
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(533)
			p.Match(AddressParserComma)
		}
		p.SetState(536)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(534)
				p.Mailbox()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(535)
				p.Cfws()
			}

		}

		p.SetState(542)
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
	p.EnterRule(localctx, 80, AddressParserRULE_obsAddrList)
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
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 82, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(544)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(543)
					p.Cfws()
				}

			}
			{
				p.SetState(546)
				p.Match(AddressParserComma)
			}

		}
		p.SetState(551)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 82, p.GetParserRuleContext())
	}
	{
		p.SetState(552)
		p.Address()
	}
	p.SetState(560)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserComma {
		{
			p.SetState(553)
			p.Match(AddressParserComma)
		}
		p.SetState(556)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(554)
				p.Address()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(555)
				p.Cfws()
			}

		}

		p.SetState(562)
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
	p.EnterRule(localctx, 82, AddressParserRULE_obsGroupList)
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
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(564)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
				{
					p.SetState(563)
					p.Cfws()
				}

			}
			{
				p.SetState(566)
				p.Match(AddressParserComma)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(569)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 86, p.GetParserRuleContext())
	}
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserTAB)|(1<<AddressParserCR)|(1<<AddressParserSP)|(1<<AddressParserLParens))) != 0 {
		{
			p.SetState(571)
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
	p.EnterRule(localctx, 84, AddressParserRULE_obsLocalPart)
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
		p.SetState(574)
		p.Word()
	}
	p.SetState(579)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserPeriod {
		{
			p.SetState(575)
			p.Match(AddressParserPeriod)
		}
		{
			p.SetState(576)
			p.Word()
		}

		p.SetState(581)
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
	p.EnterRule(localctx, 86, AddressParserRULE_obsDomain)
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
		p.SetState(582)
		p.Atom()
	}
	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AddressParserPeriod {
		{
			p.SetState(583)
			p.Match(AddressParserPeriod)
		}
		{
			p.SetState(584)
			p.Atom()
		}

		p.SetState(589)
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
	p.EnterRule(localctx, 88, AddressParserRULE_encodedWord)

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
		p.SetState(590)
		p.Match(AddressParserEqual)
	}
	{
		p.SetState(591)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(592)
		p.Charset()
	}
	{
		p.SetState(593)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(594)
		p.Encoding()
	}
	{
		p.SetState(595)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(596)
		p.EncodedText()
	}
	{
		p.SetState(597)
		p.Match(AddressParserQuestion)
	}
	{
		p.SetState(598)
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
	p.EnterRule(localctx, 90, AddressParserRULE_charset)

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
		p.SetState(600)
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
	p.EnterRule(localctx, 92, AddressParserRULE_encoding)

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
		p.SetState(602)
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
	p.EnterRule(localctx, 94, AddressParserRULE_token)
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
	p.SetState(605)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserMinus)|(1<<AddressParserDigit))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AddressParserAlphaUpper-33))|(1<<(AddressParserBackslash-33))|(1<<(AddressParserCaret-33))|(1<<(AddressParserUnderscore-33))|(1<<(AddressParserBacktick-33))|(1<<(AddressParserAlphaLower-33))|(1<<(AddressParserLCurly-33))|(1<<(AddressParserPipe-33))|(1<<(AddressParserRCurly-33))|(1<<(AddressParserTilde-33)))) != 0) {
		{
			p.SetState(604)
			p.TokenChar()
		}

		p.SetState(607)
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
	p.EnterRule(localctx, 96, AddressParserRULE_tokenChar)
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
		p.SetState(609)
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
	p.EnterRule(localctx, 98, AddressParserRULE_encodedText)
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
	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AddressParserExclamation)|(1<<AddressParserDQuote)|(1<<AddressParserHash)|(1<<AddressParserDollar)|(1<<AddressParserPercent)|(1<<AddressParserAmpersand)|(1<<AddressParserSQuote)|(1<<AddressParserLParens)|(1<<AddressParserRParens)|(1<<AddressParserAsterisk)|(1<<AddressParserPlus)|(1<<AddressParserComma)|(1<<AddressParserMinus)|(1<<AddressParserPeriod)|(1<<AddressParserSlash)|(1<<AddressParserDigit)|(1<<AddressParserColon)|(1<<AddressParserSemicolon)|(1<<AddressParserLess)|(1<<AddressParserEqual)|(1<<AddressParserGreater))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AddressParserAt-32))|(1<<(AddressParserAlphaUpper-32))|(1<<(AddressParserLBracket-32))|(1<<(AddressParserBackslash-32))|(1<<(AddressParserRBracket-32))|(1<<(AddressParserCaret-32))|(1<<(AddressParserUnderscore-32))|(1<<(AddressParserBacktick-32))|(1<<(AddressParserAlphaLower-32))|(1<<(AddressParserLCurly-32))|(1<<(AddressParserPipe-32))|(1<<(AddressParserRCurly-32))|(1<<(AddressParserTilde-32)))) != 0) {
		{
			p.SetState(611)
			p.EncodedChar()
		}

		p.SetState(614)
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
	p.EnterRule(localctx, 100, AddressParserRULE_encodedChar)
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
		p.SetState(616)
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
	p.EnterRule(localctx, 102, AddressParserRULE_crlf)

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
		p.SetState(618)
		p.Match(AddressParserCR)
	}
	{
		p.SetState(619)
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
	p.EnterRule(localctx, 104, AddressParserRULE_wsp)
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
		p.SetState(621)
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
	p.EnterRule(localctx, 106, AddressParserRULE_vchar)
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
		p.SetState(623)
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
