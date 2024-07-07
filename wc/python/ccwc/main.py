#!/usr/bin/env python3
# import argparse
# 
# def parse_args():
#     arg_parser = argparse.ArgumentParser(
#         prog="ccwc",
#         description='CCWC: A WC Clone',
#         epilog="Coding Challenges python wordcount")
#     arg_parser.add_argument('filename', help='Input file')
#     arg_parser.add_argument('-c', '--bytes', action="store_true", help='Count bytes')
#     args = arg_parser.parse_args()
# 
#     return {
#         "count_bytes": args.bytes,
#         "filename": args.filename,
#     }
# 
# def byte_count(file):
#     return len(file.encode('utf-8'))
# 
# 
# def main():
#     args = parse_args()
#     output_str = ""
# 
#     with open(args["filename"], 'r') as f:
#         file = f.read()
# 
#     if args["count_bytes"]:
#         bytes = byte_count(file)
#         output_str += f"{bytes} "
# 
# 
#     output_str += f"{args["filename"]}"
#     print(output_str)
# 
# if __name__ == '__main__':
#     main()

import argparse
import sys

def parse_args(args):
    parser = argparse.ArgumentParser(prog="ccwc")
    parser.add_argument('filename', help='Input file')
    parser.add_argument('-c', '--bytes', action='store_true', help='Count bytes')
    return parser.parse_args(args)

def main(args=None):
    if args is None:
        args = sys.argv[1:]
    args = parse_args(args)
    print(f"Filename: {args.filename}")
    print(f"Count bytes: {args.bytes}")

if __name__ == '__main__':
    main()
