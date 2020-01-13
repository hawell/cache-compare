import matplotlib
import codecs
matplotlib.use('TkAgg')
import matplotlib.pyplot as plt
import numpy as np
import math
import sys

x = [1, 2, 4, 5, 8, 10, 15, 20, 25, 50, 100]
y1 = [0.7097473, 0.741764, 0.7767088, 0.7886098, 0.8150762, 0.8279734, 0.8532976, 0.8719802, 0.8870674, 0.9374593, 0.9938549]
y2 = [0.6629113, 0.6972718, 0.7355805, 0.7485244, 0.7793832, 0.7941365, 0.8221124, 0.8442898, 0.8627302, 0.9233748, 0.9938021]
y3 = [0.5965353, 0.6317514, 0.6715119, 0.6857653, 0.7177623, 0.7342027, 0.7667168, 0.7918459, 0.812927,  0.8885482, 0.9826483]
y4 = [0.6003255, 0.6358793, 0.6761208, 0.6903394, 0.7232911, 0.7395891, 0.77259, 0.7981646, 0.8197023, 0.8965162, 0.9879142]
y5 = [0.6035072, 0.6390947, 0.6798435, 0.6941771, 0.7269138, 0.7434719, 0.7764309, 0.8024616, 0.8239484, 0.9017233, 0.9910723]
y6 = [0.685801, 0.718568, 0.7540829, 0.7663984, 0.7935927, 0.8076347, 0.8338863, 0.8534543, 0.8698319, 0.9260223, 0.9938546]
y7 = [0.6167798, 0.6529158, 0.6938553, 0.708232, 0.7409483, 0.757594, 0.7905309, 0.8160154, 0.8371661, 0.911831, 0.9938068]
plt.plot(x, y1, 'ko-', label='optimal')
plt.plot(x, y6, 'co-', label='goburrow')
plt.plot(x, y2, 'bo-', label='ristretto')
plt.plot(x, y7, 'ro-', label='ccache')
plt.plot(x, y5, 'mo-', label='fastcache')
plt.plot(x, y4, 'yo-', label='freecache')
plt.plot(x, y3, 'go-', label='bigcache')
plt.xlabel('cache size to domain', fontsize=18)
plt.ylabel('hit rate', fontsize=16)
plt.legend(loc='upper left')
plt.show()