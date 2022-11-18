import csv
import time

def cocktailSort(list):
  last = len(list) - 1
  while True:
    swapped = False
    for i in range (last-1,1,-1):
      if list[i] > list[i+1]:
        list[i], list[i+1] = list[i+1], list[i]
        swapped = True
      
    if swapped == False: return
    
    swapped = False
    for i in range(last-1,-1,-1):
      if list[i] > list[i+1]:
        list[i], list[i+1] = list[i+1], list[i]
        swapped = True

    if swapped == False: return


#main

sizes = [ 100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000 ]
times = []

for k in range(0,5):
  times.append([])
  print(':::::::::',k)
  for i in sizes:
    list = []
    with open('data' + str(i) + '.csv', newline='') as csvfile:
        spamreader = csv.reader(csvfile, delimiter=',')
        for row in spamreader:
          for number in row:
            if(number != ''):
              list.append(int(number))

    startTime = time.time()
    cocktailSort(list)
    endTime = time.time()
    times[k].append((endTime - startTime))

    print("finished: ",i)

times.append([])
for i in range(0,len(sizes)):
  print(i)
  k = 0.0
  for j in range(0,5): 
    print(j)
    k += times[j][i]
  times[5].append(k/5.0)


with open('cocktailSortTimesPy.csv', 'w', encoding='UTF8', newline='') as file:
  writer = csv.writer(file)
  for i in range(0,6): writer.writerow(times[i])
