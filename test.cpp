#include <iostream>
#include <string> 
#include <vector>
using namespace std;

int find_text(string text, string search, string fullWord, int start) {
    cout<<text<<" - "<<search<<" - "<<fullWord<<" - "<<start<<endl;
    vector<string> textVector;
    string word;
    for(int i = 0; i <= text.length(); i++) {
        if(text[i] == ' ' || i == text.length()) {
            textVector.push_back(word);
            word = "";
        }
        else word += text[i];
    }
    bool isFullWord = false;
    if(fullWord == "Y") isFullWord = true;
    word = "";
    int index = 0;
    for(auto i = textVector.begin(); i != textVector.end(); ++i) {
        if(index - (*i).length() >= start) {
            if(isFullWord) {
                if(*i == search) return index;
                word = *i;
            }
            else {
                string subString = "";
                int subIndex = index, k = 0;
                for(int j = 0; j < (*i).length(); j++) {
                    if((*i)[j] == search[k]) {
                        subIndex++;
                        k++;
                        subString += (*i)[j];
                    }
                    if(subString == search) return subIndex;
                }
            }   
        }
        index += (*i).length() + 1;
    }
    return -1;
}


int main() { 
    cout<<find_text("aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aa","aa","Y",1)<<endl;
    cout<<"------------------------------------------------------------------------"<<endl;

    cout<<find_text("abc d","abcd","N",0)<<endl;
    cout<<"------------------------------------------------------------------------"<<endl;

    cout<<find_text("nott not","not","Y",0)<<endl;
    cout<<"------------------------------------------------------------------------"<<endl;

    cout<<find_text("Teachers asder them kid","kid","Y",1)<<endl;
    cout<<"------------------------------------------------------------------------"<<endl;

    cout<<find_text("this is a Test","Test","Y",2)<<endl;
    cout<<"------------------------------------------------------------------------"<<endl;

    cout<<find_text("bAab aaBaBa Ab aa abbb ab ba aaba aab Ba aba aa a","Ba","N",6)<<endl;
    cout<<"------------------------------------------------------------------------"<<endl;
}