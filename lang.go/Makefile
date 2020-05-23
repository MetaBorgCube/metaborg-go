INPUT=src-gen/statix/signatures/*.stx
OUTPUT=trans/analysis/signatures.stx

.SILENT:
.PHONY: signatures

# Concatenates all Statix signatures into one file
signatures:
	awk 'FNR==1{print ""}1' $(INPUT) > $(OUTPUT)
	perl -0777 -i -pe 's/^\s*module\s*.*$$(.|\r|\n)*?^\s*(?=^\s*(signature|rules)\s*$$)//gm' $(OUTPUT)
	echo 'module analysis/signatures\n// Signatures copied from src-gen/statix/signatures\n' | cat - $(OUTPUT) > temp && mv temp $(OUTPUT)
