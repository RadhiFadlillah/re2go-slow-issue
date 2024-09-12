// Code generated by re2c 3.1, DO NOT EDIT.
package main

func re2Email2Rule(bytes []byte) int {
	var count int
	var cur, mar int
	bytes = append(bytes, byte(0)) // add terminating null
	lim := len(bytes) - 1          // lim points at the terminating null

	for {
		{
			var yych byte
			yych = bytes[cur]
			switch yych {
			case '+':
				fallthrough
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy3
			case ';':
				goto yy4
			default:
				if lim <= cur {
					goto yy24
				}
				goto yy1
			}
		yy1:
			cur += 1
		yy2:
			{
				continue
			}
		yy3:
			cur += 1
			mar = cur
			yych = bytes[cur]
			switch yych {
			case '+':
				fallthrough
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case '@', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy6
			default:
				goto yy2
			}
		yy4:
			cur += 1
			mar = cur
			yych = bytes[cur]
			switch yych {
			case 'b':
				goto yy9
			default:
				goto yy2
			}
		yy5:
			cur += 1
			yych = bytes[cur]
		yy6:
			switch yych {
			case '+':
				fallthrough
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy5
			case '@':
				goto yy8
			default:
				goto yy7
			}
		yy7:
			cur = mar
			goto yy2
		yy8:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case 0x00:
				goto yy7
			case '.':
				goto yy10
			default:
				goto yy11
			}
		yy9:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case 'a':
				goto yy12
			default:
				goto yy7
			}
		yy10:
			cur += 1
			yych = bytes[cur]
		yy11:
			switch yych {
			case '-':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy10
			case '.':
				goto yy13
			default:
				goto yy7
			}
		yy12:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case 's':
				goto yy14
			default:
				goto yy7
			}
		yy13:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy15
			default:
				goto yy7
			}
		yy14:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case 'e':
				goto yy17
			default:
				goto yy7
			}
		yy15:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy15
			default:
				goto yy16
			}
		yy16:
			{
				count += 1
				continue
			}
		yy17:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '6':
				goto yy18
			default:
				goto yy7
			}
		yy18:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '4':
				goto yy19
			default:
				goto yy7
			}
		yy19:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case ',':
				goto yy20
			default:
				goto yy7
			}
		yy20:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '+':
				fallthrough
			case '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy21
			default:
				goto yy7
			}
		yy21:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '+':
				fallthrough
			case '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy21
			case '=':
				goto yy23
			default:
				goto yy22
			}
		yy22:
			{
				continue
			}
		yy23:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '=':
				goto yy23
			default:
				goto yy22
			}
		yy24:
			{
				return count
			}
		}

	}
}
