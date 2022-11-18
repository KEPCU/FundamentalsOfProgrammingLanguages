import matplotlib.pyplot as plt

sizes = [ 
        100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 
        8000, 9000, 10000, 20000, 30000, 40000, 50000 
        ]

times = [
    [ 
        0.000084,0.005125,0.019379,0.041610,0.074329,0.116676,
        0.1687696,0.2257266,0.2976812,0.3724236,0.4619382,
        1.8438654,4.1674776,7.4450562,11.5773972 
    ],
    [ 
        0E+00,7.0618E-04,2.46516E-03,6.0231E-03,1.04571E-02,
        1.8584499999999997E-02,2.6566E-02,3.798084E-02,
        5.2525619999999995E-02,6.684128E-02,8.331876E-02,
        3.6075668E-01,8.398381E-01,1.4993191600000002E+00,
        2.36201094E+00 
    ],
    [ 
        0.0006016731262207031,0.09107704162597656,0.380082368850708,
        0.8619470119476318,1.5388020038604737,2.421519899368286,
        3.522882890701294,4.79917278289795,6.264121246337891,
        7.976968431472779,9.764828872680663,40.998408508300784,
        92.16900811195373,165.19458498954774,258.4798514842987 
    ]
]


plt.plot(sizes,list(times[0]),label="Cpp") 
plt.plot(sizes,list(times[1]),label="Go") 
plt.plot(sizes,list(times[2]),label="Python") 
plt.xlabel("number data")
plt.ylabel("time (s)")
plt.legend(loc="lower right")
plt.title("Cocktail Sort")
plt.show()

plt.plot(sizes,list(times[0]),label="Cpp") 
plt.plot(sizes,list(times[1]),label="Go") 
plt.xlabel("number data")
plt.ylabel("time (s)")
plt.legend(loc="lower right")
plt.title("Cocktail Sort")
plt.show()


times = [
    [ 
        0.000017,0.000206,0.000250,0.000330,0.000415,0.000502,0.0006348,
        0.0007302,0.0007406,0.0008184,0.0009192,0.001723,0.0024994,
        0.0032802,0.0046394 
    ],
    [ 
        0E+00,0E+00,0E+00,4.0414000000000004E-04,9.42E-06,1.0958E-04,4.255E-04,
        2.0683999999999996E-04,3.1272E-04,2.3562000000000002E-04,5.4902E-04,
        8.4048E-04,1.36154E-03,2.00754E-03,2.421E-03 
    ],
    [ 
        0.0,0.0006076335906982422,0.0012085437774658203,0.0017999172210693359,
        0.0025574684143066405,0.0029863834381103514,0.003389883041381836,
        0.004581928253173828,0.005384492874145508,0.005778026580810547,
        0.006590986251831054,0.01417369842529297,0.02115340232849121,
        0.029258871078491212,0.036231613159179686 
    ]
]


plt.plot(sizes,list(times[0]),label="Cpp") 
plt.plot(sizes,list(times[1]),label="Go") 
plt.plot(sizes,list(times[2]),label="Python") 
plt.xlabel("number data")
plt.ylabel("time (s)")
plt.legend(loc="lower right")
plt.title("Counting Sort")
plt.show()

plt.plot(sizes,list(times[0]),label="Cpp") 
plt.plot(sizes,list(times[1]),label="Go") 
plt.xlabel("number data")
plt.ylabel("time (s)")
plt.legend(loc="lower right")
plt.title("Counting Sort")
plt.show()