Instruction to test the working of uploading a text file. This application will provide a response to the top 10 most occuring words in the selected file.
This application is dependent on another service http://localhost:8000/wordcounts which provided the actual response for the selected file.

http://localhost:9000/readfile

curl --location --request POST 'http://localhost:9000/readfile' \
--form 'file=@"/C:/Path/To/FileFolder/file.txt"'