# iWork utilities

This project contains a tool for converting Pages files to HTML. The initial motivation was to recover my older Pages
files, but I got carried away and wrote a converter for the current format and the older iOS format.

It is very much a work in progress, but hopefully someone finds it useful.

# Setup process

Download and build

```bash
go get https://github.com/orcastor/iwork-converter/iwork2html
go get https://github.com/orcastor/iwork-converter/iwork2text
```

# Usage

Export path

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Run conventer

```bash
./iwork-converter infile.pages outfile.html
```

# iwork2html

Support convert to json / html format

# iwork2text

Support OCR function injection

``` go
func ConvertString(in string, ocr func(io.Reader) (string, error)) (string, error)
```

You can use [danlock/gogosseract](https://github.com/danlock/gogosseract) as backend.

``` go
var tess *gogosseract.Tesseract

func InitTesseractOCR(model string) error {
	trainingDataFile, err := os.Open(model)
	if err != nil {
		return err
	}

	ctx := context.TODO()
	cfg := gogosseract.Config{
		Language:     "eng",
		TrainingData: trainingDataFile,
	}
	// While Tesseract's logs are very useful for debugging, you have the option to silence or redirect it
	cfg.Stderr = io.Discard
	cfg.Stdout = io.Discard
	// Compile the Tesseract WASM and run it, loading in the TrainingData and setting any Config Variables provided
	tess, err = gogosseract.New(ctx, cfg)
	if err != nil {
		return err
	}
	return nil
}

func OCR(reader io.Reader) (string, error) {
	if tess == nil {
		return "", fmt.Errorf("tesseract not initialized")
	}

	ctx := context.TODO()
	err := tess.LoadImage(ctx, reader, gogosseract.LoadImageOptions{})
	if err != nil {
		return "", err
	}

	var text string
	text, err = tess.GetText(ctx, nil)
	if err != nil {
		return "", err
	}
	return text, nil
}
```

## Pages '13

I'm building on [the work of Sean Patrick O'Brien](https://github.com/obriensp/iWorkFileFormat) on github. He determined
the base format of the `.iwa` files in the iWork'13 were a snappy compressed sequence of protobuf-encoded records and wrote
[a tool to extract the protobuf definitions from the executables](https://github.com/obriensp/proto-dump). Those `.proto`
files are included in the `proto` directory. He also extracted tables of `int` to `type`, needed to decode the `.iwa`
archives. Those `.json` files are included in this project.

I've used the json files to generate some of the code in the `index` directory. (Using the code found in `codegen`.) I ran
`protoc` on the `.proto` files and cleaned up the results so they would compile.

On top of this, I wrote the `index` package, which loads the database into memory. And I wrote `iwork2html` which will load
a pages file and render the contents to HTML.

## iOS '.pages-tef' files

Before the format change in Pages'13, the iOS version of pages introduced a `.pages-tef` bundle format for iCloud storage.
It turns out that the sqlite database within these bundles mirror the '13 format. The `iwork2html` program handles these
files too.


## Pages'08 and Pages'09

I have included an XSLT file (`pages08tohtml.xsl`) that will process the xml from Pages'08 and Pages'09 files to HTML. You can apply it to the
index.xml.gz in a Pages'08 file bundle (which is a directory) or the index.xml found within a Pages'09 file
(which is just a zip file).  For now I'll leave it as an exercise to write a wrapper script.

