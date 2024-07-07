#!/usr/bin/env python3
import argparse
import os
import sys

def parse_args():
    arg_parser = argparse.ArgumentParser(
        prog="ccwc",
        description='CCWC: A WC Clone',
        epilog="Coding Challenges python wordcount")
    arg_parser.add_argument('filename', nargs="?", default=None, help='Input file')
    arg_parser.add_argument('-c', '--bytes', action="store_true", help='Count bytes')
    arg_parser.add_argument('-l', '--lines', action="store_true", help='Count lines')
    arg_parser.add_argument('-w', '--words', action="store_true", help='Count words')
    return arg_parser.parse_args()

def count_all(filename=None):
    if filename:
        bytes_count = os.path.getsize(filename)
        with open(filename, 'rb') as f:
            content = f.read()
    else:
        content = sys.stdin.buffer.read()
        bytes_count = len(content)

    lines_count = content.count(b'\n')
    words_count = len(content.split())

    return bytes_count, lines_count, words_count

def main():
    args = parse_args()
    output_str = ""

    if args.filename:
        bytes_count, lines_count, words_count = count_all(args.filename)
    else:
        bytes_count, lines_count, words_count = count_all()

    if args.lines:
        output_str += f"{lines_count} "

    if args.words:
        output_str += f"{words_count} "

    if args.bytes:
        output_str += f"{bytes_count} "

    if not any([args.lines, args.words, args.bytes]):
        print(f"{lines_count} {words_count} {bytes_count} {args.filename}")
        sys.exit(0)

    output_str += f"{args.filename}"
    print(output_str.strip())

if __name__ == '__main__':
    main()
