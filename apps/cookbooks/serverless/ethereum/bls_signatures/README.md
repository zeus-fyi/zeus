
Create a folder in the root directory called keystores, and then place your keystore json files in here

If using AWS

You'll use the makefile command zip.keys to zip this folder, and then you'll add this zipped folder as 
a layer in your lambda function.