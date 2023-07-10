import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from matplotlib import rcParams

config = {
    "font.family": 'serif',
    "font.size": 20,
    "mathtext.fontset": 'stix',
    "font.serif": ['SimSong'],
}
rcParams.update(config)


def getdata(path, col):
    datas = []
    with open(path) as data:
        next(data)
        for line in data:
            lines = line.replace("\n", "").split(",")
            datas.append(int(lines[col]) / float(1000000000))
    return datas


def figure():
    #datas1 = getdata("../file/construction for 10000W accounts/output1.csv", 1)
    datas2 = getdata("../file/construction for 10000W accounts/output2.csv", 1)

    plt.figure(figsize=(10, 7))

    x = np.arange(0, 10000, 10)
    #plt.plot(x, datas1, '', color='#4472C4', markerfacecolor='none', label='SetBalance and Commitment')
    plt.plot(x, datas2, '', color='#FFC000', markerfacecolor='none', label='Commitment')

    plt.xlabel("Account (1W)", fontsize=20, labelpad=10)
    plt.ylabel("Time (second)", fontsize=20, labelpad=10)
    plt.title("Ethereum MPT Construction Time", fontsize=24, pad=15)
    # plt.title("Ethereum MPT Construction Time", fontsize=24, pad=15, fontweight='bold')

    plt.legend(fontsize=16)
    plt.savefig("../file/construction for 10000W accounts/output3.png")
    plt.show()


if __name__ == '__main__':
    figure()

