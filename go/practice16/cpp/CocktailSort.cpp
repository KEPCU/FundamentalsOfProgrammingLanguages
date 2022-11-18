#include <iostream>
#include <vector>
#include <chrono>
#include <ctime> 
#include <sstream>
#include <fstream>
using namespace std;

void PrintVector(vector<int> list) {
  cout<<"[ ";
  for(int i = 0; i < list.size(); i++) cout<<list[i]<<" ,";
  cout<<"]"<<endl;
}

vector<int> cocktailSort(vector<int> list) {
	int last = list.size() - 1;
  while(true) {
		bool swapped = false;
		for (int i = 0; i < last; i++) {
			if (list[i] > list[i+1]) {
        int a =	list[i];
        list[i] = list[i+1];
        list[i+1] = a;
				swapped = true;
			}
		}
		if(!swapped) return list;
		
		swapped = false;
		for (int i = last - 1; i >= 0; i--) {
			if (list[i] > list[i+1]) {
				int a =	list[i];
        list[i] = list[i+1];
        list[i+1] = a;
				swapped = true;
			}
		}
		if (!swapped) return list;
  }
  return list;
}


vector<int> Read(string fileName) {
  vector<int> list;
  vector<vector<string>> content;
  vector<string> row;
  string line, word;
  fstream file(fileName, ios::in);
  if(file.is_open()) {
    while(getline(file, line)) {
      row.clear();
      stringstream str(line);
      while(getline(str, word, ',')) row.push_back(word);
      content.push_back(row);
    }
  }
  else {
    cout<<"ERROR!!"<<endl;
    return list;
  }
  
  for(int i = 0; i < content.size(); i++) {
    for(int j = 0; j < content[i].size(); j++) {
      list.push_back(stoi(content[i][j]));
    }
  }
  return list;
}


int main() {
  vector<int> sizes = { 100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000 };
  vector<vector<double>> times;
  ofstream myFile("cocktailSortTimesCpp.csv");

  for(int k = 0; k < 5; k++) {
    cout<<":::::::::"<<k<<endl;
    vector<double> subTime;
    times.push_back(subTime);
    for(int i = 0; i < sizes.size(); i++) {
      vector<int> list = Read("data" + to_string(sizes[i]) + ".csv");
      auto startTime = chrono::system_clock::now();
      list = cocktailSort(list);
      auto endTime = chrono::system_clock::now();
      std::chrono::duration<double> elapsed_seconds = endTime - startTime;
      times[k].push_back(elapsed_seconds.count());

      cout<<"finished: "<< sizes[i]<<endl;
    }

    string cad = "";
    
    for(int i = 0; i < sizes.size(); i++) cad += to_string(times[k][i])+",";
    myFile << cad + "\n";
  }

  string cad = "";
  vector<double> subTime;

  times.push_back(subTime);
  for(int i = 0; i < times.size(); i++) {
    double avg = 0.0;
    for(int j = 0; j < 5; j++) avg += times[j][i];
    times[5].push_back(avg/5.0);
  }

  for(int i = 0; i < sizes.size(); i++) cad += to_string(times[5][i])+",";
  myFile << cad;
  myFile.close();

  return 0;
}
