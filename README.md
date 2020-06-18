# go-generate-resources

This tool writes a Go source file named resource.go containing a map of all of the files in the current directory and their byte contents.
It requires one argument which is the package name for which that source file will belong to.
The purpose of this is to have required files embedded directly in the binary.
These resource files can be HTML templates, images, whatever.

## Example

```
% go get github.com/psyb0t/go-generate-resources
% cd /path/to/my/resources
% tree
.
├── icon.ico
└── templates
    ├── error.html
    └── index.html

1 directory, 3 files
% go-generate-resources main
% tree
.
├── icon.ico
├── resources.go
└── templates
    ├── error.html
    └── index.html

1 directory, 4 files
% cat resources.go
package main

var resources map[string][]byte = map[string][]byte{

	"icon.ico": []byte{
		0,0,1,0,1,0,48,48,0,0,1,0,32,0,168,37,0,0,22,0,0,0,40,0,0,0,48,0,0,0,96,0,0,0,1,0,32,0,0,0,0,0,0,36,0,0,37,22,0,0,37,22,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,4,0,0,0,4,0,0,0,2,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,24,0,0,0,24,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,2,0,0,0,4,0,0,0,4,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,4,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,0,0,23,0,0,0,17,0,0,0,169,0,0,0,169,0,0,0,17,0,0,0,23,0,0,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,4,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,3,0,0,0,0,0,0,0,0,1,1,1,0,0,0,0,35,0,0,0,95,1,1,1,151,1,1,1,195,2,2,2,224,2,2,2,242,0,0,0,252,0,0,0,255,0,0,0,255,0,0,0,252,2,2,2,242,2,2,2,224,1,1,1,195,1,1,1,151,0,0,0,95,0,0,0,35,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,3,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,3,0,0,0,1,0,0,0,0,0,0,0,28,0,0,0,117,1,1,1,202,2,2,2,250,2,2,2,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,253,0,0,0,253,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,2,2,2,255,2,2,2,250,1,1,1,202,0,0,0,117,0,0,0,28,0,0,0,0,0,0,0,1,0,0,0,3,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,3,0,0,0,2,0,0,0,0,0,0,0,38,0,0,0,159,2,2,2,245,2,2,2,255,0,0,0,255,0,0,0,255,0,0,0,251,3,3,3,251,20,20,19,252,39,39,37,253,54,54,51,254,9,9,9,254,0,0,0,255,0,0,0,255,9,9,9,254,54,54,51,254,39,39,37,253,20,20,19,252,3,3,3,251,0,0,0,251,0,0,0,255,0,0,0,255,2,2,2,255,2,2,2,245,0,0,0,159,0,0,0,38,0,0,0,0,0,0,0,2,0,0,0,3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,4,0,0,0,0,0,0,0,15,0,0,0,144,1,1,1,250,1,1,1,255,0,0,0,254,0,0,0,251,13,13,12,252,54,53,50,254,95,94,89,255,126,125,118,255,145,144,136,255,155,154,145,255,161,160,151,255,29,29,27,255,0,0,0,255,0,0,0,255,29,29,28,255,161,160,151,255,155,154,145,255,145,144,136,255,126,125,118,255,95,94,89,255,54,53,50,254,13,13,12,252,0,0,0,251,0,0,0,254,1,1,1,255,1,1,1,250,0,0,0,144,0,0,0,15,0,0,0,0,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,4,0,0,0,0,0,0,0,69,0,0,0,226,2,2,2,255,0,0,0,253,0,0,0,251,31,31,29,254,93,92,87,255,138,137,129,255,157,156,147,255,161,159,151,255,160,158,150,255,159,158,149,255,158,156,147,255,171,169,160,255,82,81,77,255,0,0,0,255,0,0,0,255,82,81,77,255,171,169,160,255,158,156,148,255,159,158,149,255,160,158,150,255,161,160,151,255,157,156,147,255,138,137,129,255,93,92,87,255,31,31,29,254,0,0,0,251,0,0,0,253,2,2,2,255,0,0,0,226,0,0,0,69,0,0,0,0,0,0,0,4,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,3,0,0,0,0,0,0,0,123,1,1,1,255,2,2,2,254,0,0,0,251,25,24,23,254,100,99,93,255,149,148,139,255,158,157,148,255,156,155,146,255,157,156,147,255,162,160,151,255,166,165,156,255,170,168,159,255,171,169,160,255,182,180,170,255,129,128,121,255,0,0,0,255,0,0,0,255,129,128,121,255,182,180,171,255,171,169,160,255,170,168,159,255,167,165,156,255,162,160,152,255,158,156,148,255,156,155,146,255,159,157,149,255,149,148,139,255,100,99,94,255,25,24,23,254,0,0,0,251,2,2,2,254,1,1,1,255,0,0,0,123,0,0,0,0,0,0,0,3,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,2,0,0,0,0,0,0,0,154,1,1,1,255,0,0,0,251,2,2,2,253,71,70,66,255,144,143,135,255,156,155,146,255,153,152,144,255,158,157,148,255,165,164,155,255,171,169,160,255,174,172,163,255,177,175,167,255,177,175,167,255,175,173,165,255,193,192,184,255,118,118,115,255,90,90,90,255,90,90,90,255,118,118,115,255,193,192,184,255,175,173,165,255,177,176,167,255,177,175,167,255,174,172,164,255,171,169,160,255,166,164,155,255,159,157,148,255,154,152,144,255,157,155,147,255,144,143,135,255,71,70,66,255,2,2,2,253,0,0,0,251,1,1,1,255,0,0,0,154,0,0,0,0,0,0,0,2,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,157,2,2,1,255,0,0,0,250,14,14,13,254,108,107,101,255,155,154,146,255,151,150,141,255,156,154,146,255,163,161,152,255,168,166,158,255,173,171,162,255,177,175,167,255,180,178,170,255,183,181,173,255,194,192,185,255,215,214,208,255,221,220,217,255,215,215,212,255,219,218,216,255,219,218,215,255,215,215,212,255,221,220,217,255,215,214,208,255,194,192,185,255,183,181,174,255,180,178,170,255,177,175,167,255,173,171,163,255,169,167,158,255,163,161,152,255,156,154,146,255,151,150,142,255,156,155,146,255,108,107,101,255,14,14,13,254,0,0,0,250,2,2,1,255,0,0,0,157,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,134,2,2,1,255,0,0,0,250,22,22,20,255,126,125,118,255,154,153,144,255,150,149,141,255,160,158,150,255,169,167,159,255,185,182,173,255,190,188,179,255,194,192,184,255,198,196,188,255,201,199,192,255,202,200,193,255,170,169,163,255,109,109,106,255,61,61,59,255,31,31,30,255,7,7,7,255,8,8,8,255,32,32,31,255,61,60,59,255,109,109,106,255,170,169,163,255,202,200,193,255,201,200,192,255,198,196,188,255,194,192,184,255,190,188,179,255,185,183,174,255,170,168,159,255,160,159,150,255,151,149,141,255,154,153,144,255,126,125,118,255,22,22,20,255,0,0,0,250,2,2,1,255,0,0,0,134,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,85,1,1,1,255,0,0,0,252,19,19,18,255,130,129,121,255,152,151,142,255,151,150,142,255,162,160,152,255,173,171,162,255,172,170,162,255,134,132,126,255,135,134,128,255,138,137,132,255,141,140,135,255,143,142,138,255,143,142,137,255,120,119,115,255,111,110,107,255,114,113,110,255,133,132,129,255,188,187,183,255,185,184,179,255,129,128,125,255,115,114,111,255,111,110,107,255,120,119,116,255,143,142,137,255,143,142,138,255,141,140,135,255,138,137,132,255,135,134,128,255,134,132,127,255,172,170,162,255,173,171,162,255,162,161,152,255,152,150,142,255,152,151,142,255,130,129,122,255,19,19,18,255,0,0,0,252,1,1,1,255,0,0,0,85,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,26,1,1,1,231,0,0,0,255,8,8,7,253,120,119,112,255,152,151,142,255,151,150,141,255,163,161,152,255,170,168,159,255,184,182,173,255,54,54,51,255,0,0,0,255,2,2,2,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,2,2,1,255,0,0,0,255,135,134,131,255,117,117,114,255,0,0,0,255,2,2,2,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,2,2,2,255,0,0,0,255,54,54,52,255,184,182,173,255,170,168,159,255,163,162,153,255,151,150,142,255,152,151,142,255,120,119,112,255,8,8,7,253,0,0,0,255,1,1,1,231,0,0,0,26,0,0,0,0,0,0,0,1,0,0,0,3,0,0,0,0,0,0,0,154,1,1,1,255,0,0,0,251,92,91,86,255,154,153,144,255,149,148,140,255,162,160,152,255,168,166,157,255,189,187,178,255,132,130,125,255,0,0,0,255,74,73,71,255,122,121,117,255,122,122,118,255,125,124,121,255,127,126,123,255,128,128,125,255,129,128,126,255,129,128,126,255,45,45,44,255,0,0,0,255,11,11,10,255,9,9,9,255,0,0,0,255,55,55,54,255,132,131,129,255,129,128,125,255,128,128,125,255,127,126,123,255,125,124,121,255,123,122,118,255,122,121,118,255,74,73,71,255,0,0,0,255,132,130,125,255,189,187,179,255,168,166,158,255,162,161,152,255,150,148,140,255,154,153,144,255,92,91,86,255,0,0,0,251,1,1,1,255,0,0,0,154,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,45,2,2,2,251,0,0,0,255,46,46,43,255,150,149,140,255,147,146,138,255,160,158,150,255,169,167,158,255,178,176,166,255,189,187,179,255,37,36,35,255,111,111,107,255,219,218,211,255,217,216,210,255,222,221,216,255,226,225,220,255,230,229,224,255,233,232,227,255,233,232,228,255,245,244,240,255,197,196,193,255,116,116,114,255,190,190,190,255,189,189,189,255,123,122,120,255,208,208,204,255,243,242,238,255,234,233,229,255,233,232,227,255,230,229,224,255,226,225,220,255,222,221,216,255,217,216,210,255,219,218,211,255,112,111,107,255,37,37,35,255,189,188,180,255,178,176,166,255,170,168,158,255,160,159,150,255,148,147,139,255,150,149,140,255,46,46,43,255,0,0,0,255,2,2,2,251,0,0,0,45,0,0,0,0,0,0,0,0,1,1,1,147,0,0,0,255,2,2,2,251,119,118,111,255,151,150,141,255,155,154,145,255,167,165,156,255,172,170,164,255,189,187,186,255,123,122,118,255,112,111,106,255,217,216,210,255,201,200,195,255,206,205,200,255,210,209,205,255,214,213,209,255,217,216,212,255,220,219,215,255,222,221,218,255,222,222,219,255,222,221,216,255,234,233,229,255,254,254,254,255,254,254,254,255,234,233,229,255,220,219,215,255,223,222,219,255,222,221,218,255,220,219,215,255,217,216,213,255,214,213,209,255,211,210,205,255,206,205,200,255,201,200,195,255,218,216,210,255,112,111,107,255,124,123,118,255,189,187,186,255,172,170,165,255,168,166,156,255,156,154,146,255,151,150,142,255,119,118,111,255,2,2,2,251,0,0,0,255,1,1,1,147,0,0,0,0,0,0,0,18,2,2,2,226,0,0,0,255,50,50,47,254,151,150,141,255,150,148,140,255,163,162,152,255,170,168,164,255,161,160,185,255,156,155,194,255,142,141,161,255,203,202,199,255,204,203,197,255,209,208,203,255,214,213,209,255,218,217,213,255,222,221,217,255,225,224,221,255,227,227,224,255,230,229,226,255,232,232,229,255,222,222,218,255,220,219,215,255,251,251,250,255,251,251,250,255,220,219,215,255,223,222,218,255,232,232,229,255,230,229,226,255,228,227,224,255,225,224,221,255,222,221,217,255,218,217,213,255,214,213,209,255,210,208,204,255,204,203,197,255,205,203,199,255,145,144,161,255,156,155,194,255,159,158,186,255,169,167,166,255,164,162,152,255,150,149,141,255,151,150,142,255,50,50,47,254,0,0,0,255,2,2,2,226,0,0,0,18,0,0,0,83,1,1,1,255,0,0,0,254,104,103,97,255,154,153,144,255,156,154,146,255,169,167,157,255,168,167,174,255,147,147,198,255,134,133,209,255,162,161,208,255,193,192,198,255,210,209,202,255,214,213,209,255,218,217,214,255,222,222,218,255,226,225,222,255,229,228,226,255,232,231,229,255,234,233,231,255,237,236,234,255,230,230,227,255,227,226,223,255,254,254,254,255,254,254,254,255,227,226,223,255,230,230,227,255,237,236,234,255,234,233,231,255,232,231,229,255,229,229,226,255,226,225,222,255,223,222,218,255,219,218,214,255,214,213,209,255,210,209,203,255,196,195,198,255,167,166,207,255,135,134,209,255,144,143,200,255,165,164,177,255,170,168,158,255,156,155,147,255,154,153,144,255,104,103,97,255,0,0,0,254,1,1,1,255,0,0,0,83,2,2,2,150,0,0,0,255,16,16,15,252,137,136,128,255,152,151,143,255,162,160,152,255,174,172,162,255,174,173,176,255,156,155,196,255,146,146,206,255,166,165,203,255,200,199,202,255,214,213,207,255,217,217,213,255,222,221,218,255,226,225,222,255,230,229,226,255,233,232,230,255,235,235,233,255,237,237,235,255,240,240,238,255,237,236,234,255,231,231,228,255,254,254,254,255,254,254,254,255,231,231,228,255,237,236,234,255,240,240,238,255,237,237,235,255,235,235,233,255,233,232,230,255,230,229,226,255,226,226,222,255,222,222,218,255,218,217,213,255,214,213,207,255,203,202,202,255,170,169,202,255,147,146,206,255,153,152,198,255,172,170,178,255,174,172,162,255,163,161,153,255,153,152,143,255,137,136,129,255,16,16,15,252,0,0,0,255,2,2,2,150,3,3,2,202,0,0,0,255,45,45,42,253,152,151,142,255,154,153,144,255,167,165,157,255,177,175,166,255,184,182,175,255,182,180,188,255,180,179,197,255,195,194,200,255,211,210,205,255,216,215,211,255,221,220,217,255,225,225,222,255,229,229,226,255,233,232,230,255,236,236,233,255,239,238,236,255,241,240,239,255,243,243,241,255,242,242,240,255,236,236,234,255,254,254,254,255,254,254,254,255,236,236,234,255,242,242,241,255,243,243,241,255,241,240,239,255,239,238,237,255,236,236,234,255,233,233,230,255,230,229,226,255,226,225,222,255,221,220,217,255,216,215,211,255,211,210,205,255,197,196,200,255,181,180,197,255,180,179,189,255,184,182,176,255,177,175,167,255,168,166,157,255,155,153,145,255,153,151,143,255,45,45,42,253,0,0,0,255,3,3,2,202,3,3,3,234,0,0,0,255,70,70,66,254,158,157,148,255,158,156,147,255,171,169,161,255,180,178,170,255,188,187,179,255,196,195,188,255,203,201,195,255,209,208,202,255,214,213,209,255,219,218,215,255,224,223,220,255,228,228,225,255,232,232,229,255,236,235,233,255,239,238,237,255,241,241,239,255,243,243,242,255,245,245,244,255,246,246,245,255,242,242,240,255,254,254,254,255,254,254,254,255,242,242,240,255,246,246,245,255,245,245,244,255,243,243,242,255,241,241,240,255,239,238,237,255,236,235,233,255,232,232,229,255,228,228,225,255,224,223,220,255,219,219,215,255,214,213,209,255,210,209,202,255,203,202,195,255,197,195,188,255,189,187,179,255,180,178,170,255,172,170,161,255,158,157,148,255,158,157,148,255,70,70,66,254,0,0,0,255,3,3,3,234,2,2,2,250,0,0,0,255,87,86,81,255,161,159,150,255,161,159,150,255,172,170,161,255,181,179,171,255,189,187,180,255,195,194,188,255,202,201,196,255,209,208,203,255,215,214,210,255,221,220,216,255,225,225,222,255,230,229,227,255,233,233,231,255,237,237,235,255,240,240,238,255,243,243,241,255,245,245,244,255,247,247,246,255,249,249,248,255,247,247,246,255,254,254,254,255,254,254,254,255,247,247,246,255,249,249,248,255,247,247,246,255,245,245,244,255,243,243,242,255,240,240,239,255,237,237,235,255,234,233,231,255,230,229,227,255,226,225,222,255,221,220,216,255,215,214,210,255,209,208,203,255,203,201,196,255,196,194,188,255,189,188,181,255,181,179,171,255,173,171,162,255,161,159,151,255,161,160,151,255,87,86,81,255,0,0,0,255,2,2,2,250,2,2,2,255,0,0,0,255,95,94,89,255,162,161,152,255,166,164,156,255,183,182,174,255,196,195,189,255,208,206,201,255,215,214,209,255,221,220,216,255,226,226,222,255,231,230,228,255,235,235,232,255,239,238,236,255,242,242,240,255,245,245,243,255,247,247,246,255,249,249,248,255,251,251,250,255,252,252,252,255,253,253,253,255,252,252,252,255,252,252,252,255,255,255,255,255,255,255,255,255,252,252,252,255,252,252,252,255,253,253,253,255,252,252,252,255,251,251,250,255,249,249,248,255,247,247,246,255,245,245,243,255,242,242,240,255,239,239,237,255,235,235,232,255,231,231,228,255,227,226,223,255,221,221,217,255,215,214,210,255,208,207,201,255,197,195,189,255,184,182,175,255,167,165,156,255,163,162,153,255,95,94,89,255,0,0,0,255,2,2,2,255,2,2,2,251,0,0,0,255,95,94,89,255,164,163,153,255,171,170,162,255,205,204,197,255,179,178,174,255,55,55,55,255,66,66,66,255,74,74,74,255,81,81,81,255,87,88,88,255,93,93,93,255,95,95,96,255,96,96,96,255,93,93,93,255,87,87,88,255,81,81,81,255,75,75,75,255,63,63,63,255,78,78,77,255,217,216,212,255,255,255,255,255,254,254,254,255,254,254,254,255,255,255,255,255,216,215,211,255,78,78,77,255,62,63,63,255,75,75,75,255,81,81,81,255,87,87,87,255,92,92,92,255,95,95,96,255,96,96,96,255,93,93,93,255,88,88,88,255,81,81,81,255,74,74,74,255,67,67,66,255,55,55,55,255,179,178,174,255,206,204,198,255,173,171,163,255,165,163,154,255,95,94,89,255,0,0,0,255,2,2,2,251,3,3,3,236,0,0,0,255,86,85,80,254,166,165,156,255,172,171,163,255,196,194,187,255,177,176,170,255,12,12,12,255,9,9,9,255,19,19,19,255,26,26,26,255,32,32,32,255,38,38,37,255,40,40,40,255,41,41,40,255,38,38,38,255,33,33,33,255,26,26,26,255,21,21,21,255,4,4,4,255,44,43,41,255,216,214,210,255,255,255,255,255,254,254,253,255,254,254,253,255,255,255,255,255,215,214,209,255,43,43,41,255,3,3,4,255,21,21,20,255,26,26,26,255,33,33,33,255,38,38,38,255,40,40,40,255,40,40,40,255,38,38,38,255,33,33,33,255,26,26,26,255,19,19,19,255,10,10,10,255,12,12,12,255,177,176,171,255,197,195,188,255,174,172,164,255,167,165,156,255,86,85,81,254,0,0,0,255,3,3,3,236,3,3,3,207,0,0,0,255,67,67,63,253,167,165,156,255,174,172,164,255,187,185,178,255,198,196,189,255,53,53,51,255,8,8,8,255,25,25,25,255,28,28,28,255,33,33,33,255,37,37,37,255,39,39,39,255,39,39,39,255,37,37,37,255,34,34,34,255,28,28,28,255,27,27,26,255,5,5,5,255,93,93,89,255,236,235,231,255,255,255,255,255,254,254,254,255,254,254,254,255,255,255,255,255,236,235,231,255,93,92,88,255,4,4,5,255,26,26,26,255,28,28,28,255,33,33,33,255,37,37,37,255,39,39,39,255,39,39,39,255,37,37,37,255,33,33,33,255,28,28,28,255,25,25,25,255,8,8,8,255,53,53,52,255,198,196,190,255,188,186,179,255,175,173,165,255,168,166,157,255,68,67,63,253,0,0,0,255,3,3,3,206,3,3,3,159,0,0,0,255,40,40,38,252,164,162,153,255,172,170,161,255,190,189,182,255,193,191,183,255,128,127,123,255,0,0,0,255,18,18,18,255,19,19,19,255,23,23,23,255,25,25,25,255,27,27,27,255,27,27,27,255,26,26,26,255,23,23,23,255,20,20,20,255,13,13,13,255,13,13,13,255,163,162,155,255,248,248,246,255,255,255,255,255,254,254,254,255,254,254,254,255,255,255,255,255,248,248,246,255,163,161,155,255,13,13,13,255,13,13,13,255,20,20,20,255,22,22,22,255,26,26,26,255,27,27,27,255,27,27,27,255,26,26,26,255,23,23,23,255,19,19,19,255,18,18,18,255,0,0,0,255,129,128,123,255,194,192,184,255,191,190,183,255,172,171,162,255,164,163,154,255,41,40,38,252,0,0,0,255,3,3,3,159,2,2,2,97,0,0,0,255,11,11,10,253,148,147,139,255,170,168,159,255,196,194,188,255,187,185,178,255,185,183,174,255,68,67,65,255,0,0,0,255,15,15,15,255,16,16,16,255,17,17,17,255,18,18,18,255,18,18,18,255,17,17,17,255,17,17,17,255,12,12,12,255,0,0,0,255,102,100,95,255,217,216,210,255,254,255,255,255,254,254,253,255,255,255,255,255,255,255,255,255,254,254,253,255,254,255,255,255,216,215,210,255,101,100,94,255,0,0,0,255,12,12,12,255,17,17,17,255,17,17,17,255,18,18,18,255,18,18,18,255,17,17,17,255,17,17,16,255,15,15,15,255,0,0,0,255,68,68,65,255,186,184,175,255,187,186,179,255,196,195,189,255,171,169,159,255,149,147,139,255,11,11,10,253,0,0,0,255,2,2,2,97,1,1,1,32,2,2,2,239,0,0,0,255,110,109,103,254,178,176,166,255,180,178,170,255,212,211,207,255,177,175,167,255,173,171,162,255,64,63,61,255,0,0,0,255,2,2,2,255,7,7,7,255,10,10,10,255,10,10,10,255,6,6,6,255,0,0,0,255,3,3,3,255,89,88,83,255,196,195,187,255,253,253,253,255,255,255,255,255,253,253,253,255,254,254,254,255,254,254,254,255,253,253,253,255,255,255,255,255,253,253,253,255,196,194,187,255,89,88,83,255,3,3,3,255,0,0,0,255,6,6,6,255,10,10,10,255,11,11,10,255,7,7,7,255,2,2,2,255,0,0,0,255,64,63,61,255,173,171,163,255,177,176,167,255,213,212,207,255,180,178,170,255,178,176,167,255,111,109,104,254,0,0,0,255,2,2,2,239,1,1,1,32,0,0,0,0,3,3,3,177,0,0,0,255,53,53,50,252,177,175,166,255,174,172,163,255,202,201,195,255,217,216,213,255,176,174,167,255,169,167,157,255,111,110,104,255,39,39,37,255,9,9,9,255,1,1,1,255,2,2,2,255,14,13,13,255,51,50,47,255,123,121,114,255,195,194,187,255,252,252,251,255,117,117,117,255,172,172,172,255,255,255,255,255,251,251,251,255,251,251,251,255,255,255,255,255,172,172,172,255,117,117,117,255,252,252,251,255,195,194,186,255,123,121,113,255,51,50,47,255,13,13,13,255,2,2,2,255,1,1,1,255,9,9,9,255,39,39,37,255,111,110,104,255,169,167,158,255,176,175,167,255,218,217,213,255,202,201,195,255,174,172,164,255,177,175,166,255,53,53,50,252,0,0,0,255,3,3,3,177,0,0,0,0,0,0,0,0,1,1,1,84,0,0,0,255,3,3,3,252,142,141,133,255,181,179,170,255,189,187,180,255,213,211,207,255,223,223,220,255,193,192,187,255,175,174,166,255,161,159,150,255,136,135,127,255,121,119,112,255,123,121,114,255,143,142,133,255,174,172,164,255,216,216,211,255,255,255,255,255,124,124,124,255,0,0,0,255,136,136,136,255,255,255,255,255,249,249,249,255,249,249,249,255,255,255,255,255,136,136,136,255,0,0,0,255,125,125,124,255,255,255,255,255,216,215,211,255,173,172,164,255,143,141,133,255,123,121,114,255,121,119,112,255,136,135,127,255,161,160,150,255,176,174,166,255,194,193,188,255,224,223,220,255,213,212,207,255,189,187,180,255,181,179,170,255,143,141,134,255,3,3,3,252,0,0,0,255,1,1,1,84,0,0,0,0,0,0,0,0,1,1,1,5,3,3,3,211,0,0,0,255,66,66,62,253,184,182,173,255,181,180,172,255,206,205,200,255,219,218,213,255,233,232,229,255,224,224,222,255,209,208,205,255,201,200,196,255,200,199,194,255,207,206,202,255,226,225,223,255,253,253,253,255,251,251,251,255,102,102,101,255,0,0,0,255,105,105,105,255,255,255,255,255,250,250,250,255,251,251,251,255,251,251,251,255,250,250,250,255,255,255,255,255,105,105,105,255,0,0,0,255,102,102,101,255,251,251,251,255,253,253,253,255,226,225,223,255,207,206,202,255,200,199,194,255,201,200,196,255,209,208,205,255,224,224,222,255,233,232,229,255,219,218,214,255,206,205,200,255,182,180,172,255,185,183,174,255,67,66,62,253,0,0,0,255,3,3,3,211,1,1,1,5,0,0,0,0,0,0,0,3,0,0,0,0,1,1,1,94,0,0,0,255,0,0,0,252,138,136,129,255,188,186,177,255,186,185,179,255,185,184,181,255,203,202,198,255,233,232,228,255,250,249,246,255,255,255,254,255,254,254,255,255,255,255,255,255,242,242,241,255,158,158,157,255,26,26,26,255,0,0,0,255,136,136,136,255,255,255,255,255,249,249,248,255,249,249,249,255,250,250,250,255,250,250,250,255,249,249,249,255,249,249,248,255,255,255,255,255,136,136,136,255,0,0,0,255,26,26,26,255,159,158,157,255,242,242,241,255,255,255,255,255,254,254,255,255,255,255,254,255,250,250,247,255,233,232,228,255,203,202,199,255,185,184,181,255,186,185,179,255,188,186,178,255,138,136,129,255,0,0,0,252,0,0,0,255,1,1,1,94,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,1,0,0,0,0,2,2,2,192,0,0,0,255,38,37,35,252,179,177,168,255,186,184,176,255,190,189,185,255,150,149,147,255,124,123,121,255,124,123,122,255,124,123,122,255,109,109,107,255,71,71,71,255,22,22,21,255,7,7,7,255,88,88,87,255,214,214,213,255,255,255,255,255,245,245,244,255,247,247,246,255,249,249,248,255,249,249,248,255,249,249,248,255,249,249,248,255,247,247,246,255,245,245,244,255,255,255,255,255,214,214,213,255,88,88,88,255,7,7,7,255,22,22,21,255,71,71,71,255,109,109,107,255,124,123,122,255,124,123,122,255,124,123,121,255,150,149,147,255,190,189,185,255,186,184,177,255,179,178,169,255,38,38,36,252,0,0,0,255,2,2,2,192,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,50,3,2,2,248,0,0,0,255,83,82,78,254,193,191,182,255,194,193,186,255,228,227,223,255,207,206,203,255,168,168,166,255,139,139,137,255,127,126,125,255,138,137,136,255,178,178,177,255,233,233,231,255,255,255,255,255,250,250,249,255,242,241,240,255,246,246,245,255,247,247,246,255,247,247,246,255,248,248,247,255,248,248,247,255,247,247,246,255,247,247,246,255,246,246,245,255,242,242,240,255,250,250,250,255,255,255,255,255,233,233,231,255,178,178,177,255,138,137,136,255,127,126,125,255,139,139,137,255,169,168,166,255,207,206,203,255,228,227,224,255,194,193,186,255,194,192,183,255,83,82,78,254,0,0,0,255,3,3,2,248,0,0,0,50,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,4,0,0,0,0,0,0,0,113,2,2,2,255,0,0,0,251,116,115,110,255,196,194,186,255,194,193,187,255,232,232,228,255,243,242,239,255,250,249,246,255,254,254,251,255,255,255,253,255,252,251,249,255,244,243,242,255,238,238,237,255,241,241,240,255,245,244,243,255,245,245,244,255,246,246,245,255,246,246,245,255,247,246,246,255,247,246,246,255,246,246,245,255,246,246,245,255,245,245,244,255,245,244,243,255,242,241,240,255,238,238,237,255,244,243,242,255,252,251,250,255,255,255,253,255,254,254,252,255,250,249,246,255,243,242,239,255,232,232,229,255,194,193,187,255,197,195,186,255,117,115,110,255,0,0,0,251,2,2,2,255,0,0,0,113,0,0,0,0,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,3,1,1,1,0,1,1,1,160,1,1,1,255,3,3,3,250,132,131,125,255,199,197,189,255,199,197,192,255,229,228,226,255,231,231,228,255,231,230,228,255,234,233,231,255,236,236,234,255,240,239,238,255,242,242,240,255,243,243,241,255,244,244,243,255,245,245,243,255,245,245,244,255,246,245,245,255,246,246,245,255,246,246,245,255,246,246,245,255,245,245,244,255,245,245,243,255,244,244,243,255,243,243,242,255,242,242,240,255,240,239,238,255,236,236,234,255,234,233,232,255,231,230,228,255,231,231,228,255,229,229,226,255,199,198,192,255,200,198,190,255,132,131,125,255,3,3,3,250,1,1,1,255,1,1,1,160,1,1,1,0,0,0,0,3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,5,5,5,1,0,0,0,5,1,1,1,185,0,0,0,255,5,5,5,250,130,128,123,255,202,200,192,255,199,198,193,255,230,230,227,255,240,239,238,255,238,237,235,255,239,239,237,255,241,241,239,255,242,242,241,255,243,243,242,255,244,244,243,255,245,244,243,255,245,245,244,255,245,245,244,255,246,245,244,255,246,245,244,255,245,245,244,255,245,245,244,255,245,244,243,255,244,244,243,255,243,243,242,255,242,242,241,255,241,241,239,255,239,239,237,255,238,237,235,255,240,239,238,255,230,230,227,255,200,198,193,255,202,200,192,255,130,129,123,255,5,5,5,250,0,0,0,255,1,1,1,185,0,0,0,5,5,5,5,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,11,1,1,1,189,1,1,1,255,1,1,1,250,108,107,102,254,200,198,190,255,200,199,193,255,222,221,218,255,242,242,241,255,243,243,242,255,242,241,240,255,242,242,241,255,244,243,242,255,245,244,243,255,245,245,244,255,246,245,244,255,246,246,245,255,246,246,245,255,246,246,245,255,246,246,245,255,246,245,244,255,245,245,244,255,245,244,243,255,244,244,242,255,242,242,241,255,242,241,240,255,243,243,242,255,242,242,241,255,222,221,218,255,201,199,193,255,200,199,191,255,108,107,103,254,1,1,1,250,1,1,1,255,1,1,1,189,0,0,0,11,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,23,23,22,0,0,0,0,7,1,1,1,172,2,2,2,255,0,0,0,251,67,66,64,253,180,178,171,255,205,203,196,255,209,208,204,255,232,231,229,255,246,246,245,255,248,248,247,255,246,246,245,255,246,245,244,255,246,245,244,255,246,246,245,255,246,246,245,255,247,246,245,255,247,246,245,255,246,246,245,255,246,246,245,255,246,245,244,255,246,245,244,255,246,246,245,255,248,248,247,255,246,246,245,255,232,232,229,255,210,209,204,255,205,204,196,255,180,179,172,255,67,66,64,253,0,0,0,251,2,2,2,255,1,1,1,172,0,0,0,7,23,23,22,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,3,2,2,2,22,22,21,0,1,1,1,129,3,3,3,255,0,0,0,255,20,20,19,251,123,122,117,254,195,193,186,255,208,207,200,255,214,213,208,255,229,229,226,255,243,243,242,255,250,250,250,255,252,252,252,255,252,252,252,255,252,252,251,255,251,251,251,255,251,251,251,255,252,252,251,255,252,252,252,255,252,252,252,255,250,250,250,255,243,243,242,255,229,229,226,255,214,213,208,255,208,207,200,255,195,194,186,255,123,122,117,254,20,20,19,251,0,0,0,255,3,3,3,255,1,1,1,129,22,22,21,0,3,3,2,2,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,4,1,1,1,0,0,0,0,65,3,3,3,220,1,1,1,255,0,0,0,253,39,39,37,251,128,127,123,253,187,186,179,255,208,206,199,255,213,212,206,255,218,218,213,255,226,225,222,255,232,232,230,255,237,237,235,255,239,239,238,255,239,239,238,255,237,237,235,255,232,232,230,255,226,225,222,255,219,218,213,255,213,212,206,255,208,207,200,255,188,186,179,255,129,128,123,253,39,39,38,251,0,0,0,253,1,1,1,255,3,3,3,220,0,0,0,65,1,1,1,0,1,1,1,4,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,4,0,0,0,0,0,0,0,10,3,3,3,126,2,2,2,241,0,0,0,255,0,0,0,254,24,24,23,252,88,87,84,251,144,143,138,253,180,178,172,254,198,197,190,255,207,206,199,255,211,210,203,255,213,211,204,255,213,211,204,255,211,210,203,255,207,206,199,255,199,197,190,255,180,178,172,254,144,143,138,253,88,87,84,251,24,24,23,252,0,0,0,254,0,0,0,255,2,2,2,241,3,3,3,126,0,0,0,10,0,0,0,0,0,0,0,4,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,3,1,1,1,2,0,0,0,0,0,0,0,23,4,4,4,129,2,2,2,227,0,0,0,255,0,0,0,254,0,0,0,255,13,13,13,255,43,43,41,252,70,69,67,251,89,88,85,251,98,97,94,251,98,97,94,251,89,88,85,251,70,69,67,251,43,43,41,252,13,13,13,255,0,0,0,255,0,0,0,254,0,0,0,255,2,2,2,227,4,4,4,129,0,0,0,23,0,0,0,0,1,1,1,2,0,0,0,3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,3,2,2,2,1,0,0,0,0,0,0,0,10,4,4,4,77,5,5,5,159,2,2,2,220,0,0,0,252,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,255,0,0,0,252,2,2,2,220,5,5,5,159,4,4,4,77,0,0,0,10,0,0,0,0,2,2,2,1,0,0,0,3,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,4,0,0,0,0,10,10,10,0,5,5,5,0,6,6,6,5,6,6,6,38,7,7,6,78,7,7,6,111,6,6,6,133,6,6,6,144,6,6,6,144,6,6,6,133,7,7,6,111,7,7,6,78,6,6,6,38,6,6,6,5,5,5,4,0,10,10,10,0,0,0,0,0,0,0,0,4,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,3,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,4,0,0,0,3,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,2,0,0,0,4,0,0,0,4,0,0,0,5,0,0,0,4,0,0,0,4,0,0,0,5,0,0,0,4,0,0,0,4,0,0,0,2,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,255,250,10,80,95,255,0,0,255,200,144,9,19,255,0,0,255,32,0,0,4,255,0,0,254,136,0,0,17,127,0,0,249,32,0,0,4,159,0,0,244,128,0,0,1,47,0,0,233,0,0,0,0,151,0,0,210,0,0,0,0,75,0,0,164,0,0,0,0,37,0,0,136,0,0,0,0,17,0,0,80,0,0,0,0,10,0,0,160,0,0,0,0,5,0,0,64,0,0,0,0,2,0,0,64,0,0,0,0,2,0,0,128,0,0,0,0,1,0,0,128,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,128,0,0,0,0,1,0,0,128,0,0,0,0,1,0,0,128,0,0,0,0,1,0,0,64,0,0,0,0,2,0,0,32,0,0,0,0,4,0,0,32,0,0,0,0,4,0,0,80,0,0,0,0,10,0,0,136,0,0,0,0,17,0,0,160,0,0,0,0,5,0,0,212,0,0,0,0,43,0,0,232,0,0,0,0,23,0,0,244,128,0,0,1,47,0,0,250,64,0,0,2,95,0,0,253,32,0,0,4,191,0,0,254,72,0,0,18,127,0,0,255,162,0,0,69,255,0,0,255,200,64,2,19,255,0,0,255,250,23,232,95,255,0,0,255,254,0,0,127,255,0,0,
	},

	"templates/error.html": []byte{
		52,48,52,32,78,111,116,32,70,111,117,110,100,10,
	},

	"templates/index.html": []byte{
		72,101,108,108,111,32,87,111,114,108,100,33,10,
	},

}
```

Just move the generated `resources.go` file in your project directory and start using it.
