#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <climits>
#include <set>
#include <cstdlib>
#include <vector>
using namespace std;

int main() {
  
  int unit = 99;
  
  vector<int> sizes = { 100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000 };

  for(int s = 0; s < sizes.size(); s++) {
    int UL = sizes[s]*1.8;
    int i = 0;
    set<int, greater<int> > Set;
    ofstream myFile("data" + to_string(sizes[s]) +".csv");
    
    while(Set.size() <= sizes[s]) {
      string cad = "";
      for(int j = i; j <= i+unit; j++) {
        int number = rand() % (UL + 1);
        int prev = Set.size();
        Set.insert(number);
        if(prev < Set.size()) cad += to_string(number)+",";
        else j--;
      } 
      myFile << cad+"\n";
      if(Set.size() == sizes[s]) break;
      i += unit;
    }
    
    cout<<(Set.size())<<endl;
  }  
  return 0;
}


