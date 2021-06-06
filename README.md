# BYRON
![alt text](https://github.com/Neufal777/Byron/blob/main/img/ezgif.com-gif-maker%20(2).gif?raw=true)

It aims to be an **aggregator of documents** and books and classify them in a way that they are accessible and useful, since there are many academic resources online but sometimes it is a bit difficult to find what you are really looking for.

At the moment there is only 1 source [**Libgen.is**], and all the tests that are being done are with this source.

The operation is quite basic, it is about **scraping** the sources and being able to extract all the possible information and it is stored (eg author, isbn, year, size, pages, publisher, title, etc. ..), which is intended is that the download urls redirect to the original download link, but at the moment the files are downloaded in case the original download urls fail at some point.

what's next?

- The plan is to find a way to filter all the documents, since there are more than **3,000,000** only in libgen, at the moment there are more than **120,000** downloaded and classified, but the plan is to find some way to download the files but **compressed** or discard the same files but from different editions, or files larger than xMB, since there are even files that weigh almost **500MB**

- Obtaining documents should also be separated from classification, since when there are more sources, and a way to know if files are being **duplicated** among them is necessary.

- Implement **concurrency** for the executions, since at the moment they are done separately, in the case shown below in the image, there are **10 simultaneous executions**, and since there are many categories and duplicate files are detected, there is no problem, and actually the 10 executions work "together" .. but it would be more logical to handle it from the program itself.

![alt text](https://github.com/Neufal777/Byron/blob/main/img/Captura%20de%20pantalla%202021-06-06%20a%20las%2013.48.53.png)

Info: The files are stored in a folder called **"Repository/"** and at the moment there are more than **120.000** files and more than **2TB**, I have not uploaded it to github for obvious reasons :)

Ps: This is a Saturday project, don't expect much :) Thanks!

**Any recommendation please let me know!**



## Installation

Just import it :)

```bash
go get github.com/Neufal777/Byron
```

## Usage

```python
func main() {

	Categories := []string{
		"machine learning",
		"artificial Inteligence",
		"neuroscience",
		"maths",
		"biology",
		"medicine",
		"mit",
		"genomic",
		"physics",
		"chemistry",
		"universe",
		"paper",
	}

	for _, c := range Categories {

		fmt.Println(chalk.Magenta.Color("processing " + c))
		core.LIBGENDownloadAll(c)

	}

}

```

## Output

Generated .json with all the data from the books and the folder "Repository/" with all the documents. **[Make sure to create the folder "Repository" before running :) ]**

