#include <iostream>
#include <vector>
#include <chrono>
#include <ctime> 
#include <sstream>
#include <fstream>
#include <algorithm>
using namespace std;

void PrintVector(vector<int> list) {
  cout<<"[ ";
  for(int i = 0; i < list.size(); i++) cout<<list[i]<<" ,";
  cout<<"]"<<endl;
}

vector<int> CountingSort(vector<int> list) {
  int minItem = *min_element(list.begin(), list.end());
  vector<int> countList ((*max_element(list.begin(), list.end()) - minItem + 1), 0);
  vector<int> outputList (list.size(),0);

  for(int i = 0; i < list.size(); i++) countList[list[i] - minItem] += 1;

  for(int i = 1; i < countList.size(); i++) countList[i] += countList[i-1];

  for (int i = list.size() - 1; i >= 0; i--) {
    outputList[countList[list[i] - minItem] - 1] = list[i];
    countList[list[i] - minItem] -= 1;
  }

  for(int i = 0; i < list.size(); i++) list[i] = outputList[i];

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
  ofstream myFile("countingSortTimesCpp.csv");

  for(int k = 0; k < 5; k++) {
    cout<<":::::::::"<<k<<endl;
    vector<double> subTime;
    times.push_back(subTime);
    for(int i = 0; i < sizes.size(); i++) {
      vector<int> list = Read("data" + to_string(sizes[i]) + ".csv");
      auto startTime = chrono::system_clock::now();
      list = CountingSort(list);
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
