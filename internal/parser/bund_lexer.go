// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 376,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 5, 8, 97, 10, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 7, 9, 103, 10, 9, 12, 9, 14, 9, 106, 11, 9, 3, 9, 6, 9,
	109, 10, 9, 13, 9, 14, 9, 110, 5, 9, 113, 10, 9, 3, 10, 3, 10, 3, 10, 6,
	10, 118, 10, 10, 13, 10, 14, 10, 119, 3, 11, 3, 11, 3, 11, 6, 11, 125,
	10, 11, 13, 11, 14, 11, 126, 3, 12, 3, 12, 3, 12, 6, 12, 132, 10, 12, 13,
	12, 14, 12, 133, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 140, 10, 14, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 152,
	10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 5, 16, 166, 10, 16, 3, 17, 3, 17, 7, 17, 170, 10,
	17, 12, 17, 14, 17, 173, 11, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 5, 25, 194, 10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	27, 3, 27, 3, 28, 5, 28, 203, 10, 28, 3, 28, 6, 28, 206, 10, 28, 13, 28,
	14, 28, 207, 3, 28, 5, 28, 211, 10, 28, 3, 28, 3, 28, 5, 28, 215, 10, 28,
	3, 28, 6, 28, 218, 10, 28, 13, 28, 14, 28, 219, 3, 28, 5, 28, 223, 10,
	28, 3, 28, 5, 28, 226, 10, 28, 3, 29, 7, 29, 229, 10, 29, 12, 29, 14, 29,
	232, 11, 29, 3, 29, 3, 29, 6, 29, 236, 10, 29, 13, 29, 14, 29, 237, 3,
	29, 6, 29, 241, 10, 29, 13, 29, 14, 29, 242, 3, 29, 5, 29, 246, 10, 29,
	3, 30, 3, 30, 3, 31, 5, 31, 251, 10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3,
	32, 3, 32, 5, 32, 259, 10, 32, 3, 32, 7, 32, 262, 10, 32, 12, 32, 14, 32,
	265, 11, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 272, 10, 32, 3,
	32, 7, 32, 275, 10, 32, 12, 32, 14, 32, 278, 11, 32, 3, 32, 5, 32, 281,
	10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 288, 10, 33, 12, 33,
	14, 33, 291, 11, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 7, 33, 301, 10, 33, 12, 33, 14, 33, 304, 11, 33, 3, 33, 3, 33, 3, 33,
	5, 33, 309, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 315, 10, 34, 5,
	34, 317, 10, 34, 3, 35, 3, 35, 5, 35, 321, 10, 35, 3, 35, 5, 35, 324, 10,
	35, 3, 36, 3, 36, 3, 36, 5, 36, 329, 10, 36, 3, 37, 3, 37, 5, 37, 333,
	10, 37, 3, 37, 5, 37, 336, 10, 37, 3, 37, 3, 37, 5, 37, 340, 10, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 7, 38, 346, 10, 38, 12, 38, 14, 38, 349, 11, 38,
	3, 39, 3, 39, 3, 39, 3, 39, 7, 39, 355, 10, 39, 12, 39, 14, 39, 358, 11,
	39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 7, 40, 367, 10, 40,
	12, 40, 14, 40, 370, 11, 40, 3, 41, 6, 41, 373, 10, 41, 13, 41, 14, 41,
	374, 5, 289, 302, 356, 2, 42, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51,
	2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2, 69, 2, 71, 2,
	73, 2, 75, 2, 77, 2, 79, 2, 81, 2, 3, 2, 22, 4, 2, 81, 81, 113, 113, 3,
	2, 50, 57, 4, 2, 90, 90, 122, 122, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2,
	68, 68, 100, 100, 3, 2, 50, 51, 6, 2, 44, 45, 47, 47, 98, 98, 128, 128,
	6, 2, 38, 40, 46, 46, 62, 64, 66, 66, 3, 2, 51, 59, 3, 2, 50, 59, 4, 2,
	71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 6, 2, 12, 12, 15, 15, 41, 41, 94,
	94, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 3, 2, 94, 94, 4, 2, 67, 92, 99,
	124, 3, 2, 99, 124, 4, 2, 12, 12, 14, 15, 4, 2, 12, 12, 15, 15, 4, 2, 11,
	11, 34, 34, 2, 411, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2,
	39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2,
	2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 3, 83, 3, 2, 2, 2, 5, 85, 3, 2, 2,
	2, 7, 87, 3, 2, 2, 2, 9, 89, 3, 2, 2, 2, 11, 91, 3, 2, 2, 2, 13, 93, 3,
	2, 2, 2, 15, 96, 3, 2, 2, 2, 17, 112, 3, 2, 2, 2, 19, 114, 3, 2, 2, 2,
	21, 121, 3, 2, 2, 2, 23, 128, 3, 2, 2, 2, 25, 135, 3, 2, 2, 2, 27, 139,
	3, 2, 2, 2, 29, 151, 3, 2, 2, 2, 31, 165, 3, 2, 2, 2, 33, 167, 3, 2, 2,
	2, 35, 174, 3, 2, 2, 2, 37, 176, 3, 2, 2, 2, 39, 178, 3, 2, 2, 2, 41, 180,
	3, 2, 2, 2, 43, 182, 3, 2, 2, 2, 45, 184, 3, 2, 2, 2, 47, 186, 3, 2, 2,
	2, 49, 193, 3, 2, 2, 2, 51, 197, 3, 2, 2, 2, 53, 199, 3, 2, 2, 2, 55, 225,
	3, 2, 2, 2, 57, 245, 3, 2, 2, 2, 59, 247, 3, 2, 2, 2, 61, 250, 3, 2, 2,
	2, 63, 280, 3, 2, 2, 2, 65, 308, 3, 2, 2, 2, 67, 316, 3, 2, 2, 2, 69, 323,
	3, 2, 2, 2, 71, 328, 3, 2, 2, 2, 73, 330, 3, 2, 2, 2, 75, 341, 3, 2, 2,
	2, 77, 350, 3, 2, 2, 2, 79, 362, 3, 2, 2, 2, 81, 372, 3, 2, 2, 2, 83, 84,
	7, 126, 2, 2, 84, 4, 3, 2, 2, 2, 85, 86, 7, 42, 2, 2, 86, 6, 3, 2, 2, 2,
	87, 88, 7, 43, 2, 2, 88, 8, 3, 2, 2, 2, 89, 90, 7, 93, 2, 2, 90, 10, 3,
	2, 2, 2, 91, 92, 7, 95, 2, 2, 92, 12, 3, 2, 2, 2, 93, 94, 7, 48, 2, 2,
	94, 14, 3, 2, 2, 2, 95, 97, 5, 59, 30, 2, 96, 95, 3, 2, 2, 2, 96, 97, 3,
	2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 5, 17, 9, 2, 99, 16, 3, 2, 2, 2, 100,
	104, 5, 51, 26, 2, 101, 103, 5, 53, 27, 2, 102, 101, 3, 2, 2, 2, 103, 106,
	3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 113, 3, 2,
	2, 2, 106, 104, 3, 2, 2, 2, 107, 109, 7, 50, 2, 2, 108, 107, 3, 2, 2, 2,
	109, 110, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	113, 3, 2, 2, 2, 112, 100, 3, 2, 2, 2, 112, 108, 3, 2, 2, 2, 113, 18, 3,
	2, 2, 2, 114, 115, 7, 50, 2, 2, 115, 117, 9, 2, 2, 2, 116, 118, 9, 3, 2,
	2, 117, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119,
	120, 3, 2, 2, 2, 120, 20, 3, 2, 2, 2, 121, 122, 7, 50, 2, 2, 122, 124,
	9, 4, 2, 2, 123, 125, 9, 5, 2, 2, 124, 123, 3, 2, 2, 2, 125, 126, 3, 2,
	2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 22, 3, 2, 2, 2,
	128, 129, 7, 50, 2, 2, 129, 131, 9, 6, 2, 2, 130, 132, 9, 7, 2, 2, 131,
	130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134,
	3, 2, 2, 2, 134, 24, 3, 2, 2, 2, 135, 136, 5, 55, 28, 2, 136, 26, 3, 2,
	2, 2, 137, 140, 5, 63, 32, 2, 138, 140, 5, 65, 33, 2, 139, 137, 3, 2, 2,
	2, 139, 138, 3, 2, 2, 2, 140, 28, 3, 2, 2, 2, 141, 142, 7, 118, 2, 2, 142,
	143, 7, 116, 2, 2, 143, 144, 7, 119, 2, 2, 144, 152, 7, 103, 2, 2, 145,
	146, 7, 86, 2, 2, 146, 147, 7, 84, 2, 2, 147, 148, 7, 87, 2, 2, 148, 152,
	7, 71, 2, 2, 149, 150, 7, 37, 2, 2, 150, 152, 7, 86, 2, 2, 151, 141, 3,
	2, 2, 2, 151, 145, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152, 30, 3, 2, 2,
	2, 153, 154, 7, 104, 2, 2, 154, 155, 7, 99, 2, 2, 155, 156, 7, 110, 2,
	2, 156, 157, 7, 117, 2, 2, 157, 166, 7, 103, 2, 2, 158, 159, 7, 72, 2,
	2, 159, 160, 7, 67, 2, 2, 160, 161, 7, 78, 2, 2, 161, 162, 7, 85, 2, 2,
	162, 166, 7, 71, 2, 2, 163, 164, 7, 37, 2, 2, 164, 166, 7, 72, 2, 2, 165,
	153, 3, 2, 2, 2, 165, 158, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 32, 3,
	2, 2, 2, 167, 171, 5, 69, 35, 2, 168, 170, 5, 71, 36, 2, 169, 168, 3, 2,
	2, 2, 170, 173, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2,
	172, 34, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 174, 175, 7, 60, 2, 2, 175,
	36, 3, 2, 2, 2, 176, 177, 7, 61, 2, 2, 177, 38, 3, 2, 2, 2, 178, 179, 7,
	49, 2, 2, 179, 40, 3, 2, 2, 2, 180, 181, 9, 8, 2, 2, 181, 42, 3, 2, 2,
	2, 182, 183, 9, 9, 2, 2, 183, 44, 3, 2, 2, 2, 184, 185, 7, 96, 2, 2, 185,
	46, 3, 2, 2, 2, 186, 187, 7, 97, 2, 2, 187, 48, 3, 2, 2, 2, 188, 194, 5,
	81, 41, 2, 189, 194, 5, 75, 38, 2, 190, 194, 5, 73, 37, 2, 191, 194, 5,
	79, 40, 2, 192, 194, 5, 77, 39, 2, 193, 188, 3, 2, 2, 2, 193, 189, 3, 2,
	2, 2, 193, 190, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 193, 192, 3, 2, 2, 2,
	194, 195, 3, 2, 2, 2, 195, 196, 8, 25, 2, 2, 196, 50, 3, 2, 2, 2, 197,
	198, 9, 10, 2, 2, 198, 52, 3, 2, 2, 2, 199, 200, 9, 11, 2, 2, 200, 54,
	3, 2, 2, 2, 201, 203, 5, 59, 30, 2, 202, 201, 3, 2, 2, 2, 202, 203, 3,
	2, 2, 2, 203, 210, 3, 2, 2, 2, 204, 206, 9, 11, 2, 2, 205, 204, 3, 2, 2,
	2, 206, 207, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208,
	211, 3, 2, 2, 2, 209, 211, 5, 57, 29, 2, 210, 205, 3, 2, 2, 2, 210, 209,
	3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 214, 9, 12, 2, 2, 213, 215, 9, 13,
	2, 2, 214, 213, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 217, 3, 2, 2, 2,
	216, 218, 9, 11, 2, 2, 217, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219,
	217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 226, 3, 2, 2, 2, 221, 223,
	5, 59, 30, 2, 222, 221, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 224, 3,
	2, 2, 2, 224, 226, 5, 57, 29, 2, 225, 202, 3, 2, 2, 2, 225, 222, 3, 2,
	2, 2, 226, 56, 3, 2, 2, 2, 227, 229, 9, 11, 2, 2, 228, 227, 3, 2, 2, 2,
	229, 232, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231,
	233, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 233, 235, 7, 48, 2, 2, 234, 236,
	9, 11, 2, 2, 235, 234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 235, 3, 2,
	2, 2, 237, 238, 3, 2, 2, 2, 238, 246, 3, 2, 2, 2, 239, 241, 9, 11, 2, 2,
	240, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242,
	243, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 246, 7, 48, 2, 2, 245, 230,
	3, 2, 2, 2, 245, 240, 3, 2, 2, 2, 246, 58, 3, 2, 2, 2, 247, 248, 9, 13,
	2, 2, 248, 60, 3, 2, 2, 2, 249, 251, 7, 15, 2, 2, 250, 249, 3, 2, 2, 2,
	250, 251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 7, 12, 2, 2, 253,
	62, 3, 2, 2, 2, 254, 263, 7, 41, 2, 2, 255, 258, 7, 94, 2, 2, 256, 259,
	5, 61, 31, 2, 257, 259, 11, 2, 2, 2, 258, 256, 3, 2, 2, 2, 258, 257, 3,
	2, 2, 2, 259, 262, 3, 2, 2, 2, 260, 262, 10, 14, 2, 2, 261, 255, 3, 2,
	2, 2, 261, 260, 3, 2, 2, 2, 262, 265, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2,
	263, 264, 3, 2, 2, 2, 264, 266, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 266,
	281, 7, 41, 2, 2, 267, 276, 7, 36, 2, 2, 268, 271, 7, 94, 2, 2, 269, 272,
	5, 61, 31, 2, 270, 272, 11, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 270, 3,
	2, 2, 2, 272, 275, 3, 2, 2, 2, 273, 275, 10, 15, 2, 2, 274, 268, 3, 2,
	2, 2, 274, 273, 3, 2, 2, 2, 275, 278, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2,
	276, 277, 3, 2, 2, 2, 277, 279, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 279,
	281, 7, 36, 2, 2, 280, 254, 3, 2, 2, 2, 280, 267, 3, 2, 2, 2, 281, 64,
	3, 2, 2, 2, 282, 283, 7, 41, 2, 2, 283, 284, 7, 41, 2, 2, 284, 285, 7,
	41, 2, 2, 285, 289, 3, 2, 2, 2, 286, 288, 5, 67, 34, 2, 287, 286, 3, 2,
	2, 2, 288, 291, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2,
	290, 292, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 292, 293, 7, 41, 2, 2, 293,
	294, 7, 41, 2, 2, 294, 309, 7, 41, 2, 2, 295, 296, 7, 36, 2, 2, 296, 297,
	7, 36, 2, 2, 297, 298, 7, 36, 2, 2, 298, 302, 3, 2, 2, 2, 299, 301, 5,
	67, 34, 2, 300, 299, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302, 303, 3, 2,
	2, 2, 302, 300, 3, 2, 2, 2, 303, 305, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2,
	305, 306, 7, 36, 2, 2, 306, 307, 7, 36, 2, 2, 307, 309, 7, 36, 2, 2, 308,
	282, 3, 2, 2, 2, 308, 295, 3, 2, 2, 2, 309, 66, 3, 2, 2, 2, 310, 317, 10,
	16, 2, 2, 311, 314, 7, 94, 2, 2, 312, 315, 5, 61, 31, 2, 313, 315, 11,
	2, 2, 2, 314, 312, 3, 2, 2, 2, 314, 313, 3, 2, 2, 2, 315, 317, 3, 2, 2,
	2, 316, 310, 3, 2, 2, 2, 316, 311, 3, 2, 2, 2, 317, 68, 3, 2, 2, 2, 318,
	321, 9, 17, 2, 2, 319, 321, 5, 39, 20, 2, 320, 318, 3, 2, 2, 2, 320, 319,
	3, 2, 2, 2, 321, 324, 3, 2, 2, 2, 322, 324, 9, 18, 2, 2, 323, 320, 3, 2,
	2, 2, 323, 322, 3, 2, 2, 2, 324, 70, 3, 2, 2, 2, 325, 329, 5, 69, 35, 2,
	326, 329, 9, 11, 2, 2, 327, 329, 5, 39, 20, 2, 328, 325, 3, 2, 2, 2, 328,
	326, 3, 2, 2, 2, 328, 327, 3, 2, 2, 2, 329, 72, 3, 2, 2, 2, 330, 332, 7,
	94, 2, 2, 331, 333, 5, 81, 41, 2, 332, 331, 3, 2, 2, 2, 332, 333, 3, 2,
	2, 2, 333, 339, 3, 2, 2, 2, 334, 336, 7, 15, 2, 2, 335, 334, 3, 2, 2, 2,
	335, 336, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 340, 7, 12, 2, 2, 338,
	340, 4, 14, 15, 2, 339, 335, 3, 2, 2, 2, 339, 338, 3, 2, 2, 2, 340, 74,
	3, 2, 2, 2, 341, 342, 7, 37, 2, 2, 342, 343, 7, 37, 2, 2, 343, 347, 3,
	2, 2, 2, 344, 346, 10, 19, 2, 2, 345, 344, 3, 2, 2, 2, 346, 349, 3, 2,
	2, 2, 347, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 76, 3, 2, 2, 2,
	349, 347, 3, 2, 2, 2, 350, 351, 7, 49, 2, 2, 351, 352, 7, 44, 2, 2, 352,
	356, 3, 2, 2, 2, 353, 355, 11, 2, 2, 2, 354, 353, 3, 2, 2, 2, 355, 358,
	3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 357, 359, 3, 2,
	2, 2, 358, 356, 3, 2, 2, 2, 359, 360, 7, 44, 2, 2, 360, 361, 7, 49, 2,
	2, 361, 78, 3, 2, 2, 2, 362, 363, 7, 49, 2, 2, 363, 364, 7, 49, 2, 2, 364,
	368, 3, 2, 2, 2, 365, 367, 10, 20, 2, 2, 366, 365, 3, 2, 2, 2, 367, 370,
	3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 80, 3, 2,
	2, 2, 370, 368, 3, 2, 2, 2, 371, 373, 9, 21, 2, 2, 372, 371, 3, 2, 2, 2,
	373, 374, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375,
	82, 3, 2, 2, 2, 49, 2, 96, 104, 110, 112, 119, 126, 133, 139, 151, 165,
	171, 193, 202, 207, 210, 214, 219, 222, 225, 230, 237, 242, 245, 250, 258,
	261, 263, 271, 274, 276, 280, 289, 302, 308, 314, 316, 320, 323, 328, 332,
	335, 339, 347, 356, 368, 374, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'|'", "'('", "')'", "'['", "']'", "'.'", "", "", "", "", "", "", "",
	"", "", "", "':'", "';'", "'/'", "", "", "'^'", "'_'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "INTEGER", "DECIMAL_INTEGER", "OCT_INTEGER",
	"HEX_INTEGER", "BIN_INTEGER", "FLOAT_NUMBER", "STRING", "TRUE", "FALSE",
	"NAME", "TOBEGIN", "TOEND", "SLASH", "DIR", "CMD", "DUPLICATE", "DROP",
	"SKIP_",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "INTEGER", "DECIMAL_INTEGER",
	"OCT_INTEGER", "HEX_INTEGER", "BIN_INTEGER", "FLOAT_NUMBER", "STRING",
	"TRUE", "FALSE", "NAME", "TOBEGIN", "TOEND", "SLASH", "DIR", "CMD", "DUPLICATE",
	"DROP", "SKIP_", "NON_ZERO_DIGIT", "DIGIT", "EXPONENT_OR_POINT_FLOAT",
	"POINT_FLOAT", "SIGN", "RN", "SHORT_STRING", "LONG_STRING", "LONG_STRING_ITEM",
	"ID_START", "ID_CONTINUE", "LINE_JOINING", "COMMENT", "BLOCK_COMMENT",
	"CCOMMENT", "SPACES",
}

type BundLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewBundLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *BundLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundLexer(input antlr.CharStream) *BundLexer {
	l := new(BundLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Bund.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BundLexer tokens.
const (
	BundLexerT__0            = 1
	BundLexerT__1            = 2
	BundLexerT__2            = 3
	BundLexerT__3            = 4
	BundLexerT__4            = 5
	BundLexerT__5            = 6
	BundLexerINTEGER         = 7
	BundLexerDECIMAL_INTEGER = 8
	BundLexerOCT_INTEGER     = 9
	BundLexerHEX_INTEGER     = 10
	BundLexerBIN_INTEGER     = 11
	BundLexerFLOAT_NUMBER    = 12
	BundLexerSTRING          = 13
	BundLexerTRUE            = 14
	BundLexerFALSE           = 15
	BundLexerNAME            = 16
	BundLexerTOBEGIN         = 17
	BundLexerTOEND           = 18
	BundLexerSLASH           = 19
	BundLexerDIR             = 20
	BundLexerCMD             = 21
	BundLexerDUPLICATE       = 22
	BundLexerDROP            = 23
	BundLexerSKIP_           = 24
)
