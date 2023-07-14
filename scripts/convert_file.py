import os
import glob
from pdf2image import convert_from_path
import argparse
import json
import re
import subprocess
import msoffice2pdf

print(os.curdir)

parser = argparse.ArgumentParser(description='Convert docx to png')
parser.add_argument('--input_file', type=str, required=True, help='input file')

args = parser.parse_args()

input_file = args.input_file
output_dir = 'tmp'

docx_file = input_file

class LibreOfficeError(Exception):
    def __init__(self, output):
        self.output = output

def convert_docx_to_pdf(docx_file, output_dir):
    # 将.docx文件转换为.pdf文件
    docx_file = os.path.abspath(docx_file)
    args = ['libreoffice', '--headless', '--convert-to', 'pdf', '--outdir', output_dir, docx_file]
    process = subprocess.run(args, stdout=subprocess.PIPE, stderr=subprocess.PIPE, timeout=3000)
    filename = re.search('-> (.*?) using filter', process.stdout.decode())
    if filename is None:
        raise LibreOfficeError(process.stderr.decode())
    return os.path.join(output_dir, filename.group(1))


def convert_pdf_to_png(pdf_file, output_dir):
    # 将.pdf文件转换为.png文件
    images = convert_from_path(pdf_file)
    image_files = []
    for i, image in enumerate(images):
        image_file = os.path.join(output_dir, f'{os.path.splitext(os.path.basename(docx_file))[0]}_{i}.png')
        image.save(image_file, 'PNG')
        image_files.append(image_file)
    return image_files
    
def parse_file_ext(file):
    return os.path.splitext(file)[1]



image_files = []

if parse_file_ext(input_file) in ['.docx', '.doc']:
    # pdf_file = convert_docx_to_pdf(input_file, output_dir)
    pdf_file = msoffice2pdf.convert(input_file, output_dir, soft=1)
    image_files = convert_pdf_to_png(pdf_file, output_dir)
    os.remove(pdf_file)
elif parse_file_ext(input_file) in ['.pdf']:
    image_files = convert_pdf_to_png(input_file, output_dir)
elif parse_file_ext(input_file) in ['.png', '.jpg', '.jpeg']:
    image_files = [input_file]
else:
    raise Exception('不支持的文件格式')

json.dump(image_files, open('tmp/image_files.json', 'w'))