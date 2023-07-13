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


def getdata1(path, col):
    datas = []
    with open(path) as data:
        next(data)
        for line in data:
            lines = line.replace("\n", "").split(",")
            datas.append(int(lines[col]) / float(1000000000))
    return datas


def getdata2(path, col):
    datas = []
    with open(path) as data:
        for line in data:
            lines = line.replace("\n", "").split(",")
            datas.append(int(lines[col]))
    return datas


def figure1():
    # datas1 = getdata("../file/construction for 10000W accounts/output1.csv", 1)
    datas2 = getdata1("../file/construction for 10000W accounts/output2.csv", 1)

    plt.figure(figsize=(10, 7))

    x = np.arange(0, 10000, 10)
    # plt.plot(x, datas1, '', color='#4472C4', markerfacecolor='none', label='SetBalance and Commitment')
    plt.plot(x, datas2, '', color='#FFC000', markerfacecolor='none', label='Commitment')

    plt.xlabel("Account (W)", fontsize=20, labelpad=10)
    plt.ylabel("Time (second)", fontsize=20, labelpad=10)
    plt.title("Ethereum MPT Construction Time", fontsize=24, pad=15)
    # plt.title("Ethereum MPT Construction Time", fontsize=24, pad=15, fontweight='bold')

    plt.legend(fontsize=16)
    plt.savefig("../file/construction for 10000W accounts/output3.png")
    plt.show()


def figure2():
    name = "random read"
    file = name + " result"

    # 数据
    categories = ['1', '10', '100', '283', '1000', '100000']
    values0 = getdata2("../file/read and write/" + file + ".csv", 0)
    values1 = getdata2("../file/read and write/" + file + ".csv", 1)

    plt.figure(figsize=(15, 10))

    bar_width = 0.4
    x = np.arange(len(categories))

    plt.bar(x - bar_width/2, values0, width=bar_width, label='GetBalance without cache')
    plt.bar(x + bar_width/2, values1, width=bar_width, label='GetBalance with cache')

    plt.xlabel("Account (W)", fontsize=20, labelpad=10)
    plt.ylabel("Time (ns)", fontsize=20, labelpad=10)
    plt.title("Ethereum MPT " + name + " time", fontsize=24, pad=15)

    plt.xticks(x, categories)
    plt.legend(fontsize=16)

    # 在每个柱子顶部添加数据标签
    for i, val0 in enumerate(values0):
        plt.text(x[i] - bar_width / 2, val0 + 1, str(val0), ha='center', va='bottom', fontsize=16)

    for i, val1 in enumerate(values1):
        plt.text(x[i] + bar_width / 2, val1 + 1, str(val1), ha='center', va='bottom', fontsize=16)

    plt.savefig("../file/read and write/" + file + ".png")
    plt.savefig("../file/read and write/" + file + ".eps")
    plt.show()


def figure3():
    name = "sequential write"
    file = name + " result"

    # 数据
    categories = ['1', '10', '100', '283', '1000', '100000']
    values0 = getdata2("../file/read and write/" + file + ".csv", 0)
    values1 = getdata2("../file/read and write/" + file + ".csv", 1)
    values2 = getdata2("../file/read and write/" + file + ".csv", 2)
    values3 = getdata2("../file/read and write/" + file + ".csv", 3)
    values4 = getdata2("../file/read and write/" + file + ".csv", 4)
    values5 = getdata2("../file/read and write/" + file + ".csv", 5)

    plt.figure(figsize=(15, 10))

    bar_width = 0.1
    x = np.arange(len(categories))

    plt.bar(x - bar_width * 2, values0, width=bar_width, label='SetBalance without cache')
    plt.bar(x - bar_width, values1, width=bar_width, label='Commit to memory without cache')
    plt.bar(x, values2, width=bar_width, label='Commit to storage without cache')
    plt.bar(x + bar_width, values3, width=bar_width, label='SetBalance with cache')
    plt.bar(x + bar_width * 2, values4, width=bar_width, label='Commit to memory with cache')
    plt.bar(x + bar_width * 3, values5, width=bar_width, label='Commit to storage with cache')

    plt.xlabel("Account (W)", fontsize=20, labelpad=10)
    plt.ylabel("Time (ns)", fontsize=20, labelpad=10)
    plt.title("Ethereum MPT " + name + " time", fontsize=24, pad=15)

    plt.xticks(x, categories)
    plt.legend(fontsize=16)

    # 在每个柱子顶部添加数据标签
    # for i, val0 in enumerate(values0):
    #     plt.text(x[i] - bar_width * 2, val0 + 1, str(val0), ha='center', va='bottom', fontsize=12)
    #
    # for i, val1 in enumerate(values1):
    #     plt.text(x[i] - bar_width, val1 + 1, str(val1), ha='center', va='bottom', fontsize=12)
    #
    # for i, val2 in enumerate(values2):
    #     plt.text(x[i], val2 + 1, str(val2), ha='center', va='bottom', fontsize=12)
    #
    # for i, val3 in enumerate(values3):
    #     plt.text(x[i] + bar_width, val3 + 1, str(val3), ha='center', va='bottom', fontsize=12)
    #
    # for i, val4 in enumerate(values4):
    #     plt.text(x[i] + bar_width * 2, val4 + 1, str(val4), ha='center', va='bottom', fontsize=12)
    #
    # for i, val5 in enumerate(values5):
    #     plt.text(x[i] + bar_width * 3, val5 + 1, str(val5), ha='center', va='bottom', fontsize=12)

    plt.savefig("../file/read and write/" + file + ".png")
    plt.savefig("../file/read and write/" + file + ".eps")
    plt.show()


if __name__ == '__main__':
    figure3()
