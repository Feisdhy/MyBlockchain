import pandas as pd
import matplotlib.pyplot as plt

import pandas as pd
import matplotlib.pyplot as plt


def csv1():
    # 读取CSV文件
    df = pd.read_csv('file/output_1.csv')

    # 提取x轴和y轴数据
    x = df['accounts']
    y = df['time']

    # 绘制图形
    plt.plot(x, y)
    plt.xlabel('accumulated accounts')
    plt.ylabel('accumulated time/s')
    plt.title('Ethereum StateDB insertion for 10000W accounts')
    plt.show()
    plt.savefig('file/output_1.png')


def csv2():
    # 读取CSV文件
    df = pd.read_csv('file/output_2.csv')

    # 提取x轴和y轴数据
    x = df['accounts']
    y = df['time']

    # 绘制图形
    plt.plot(x, y)
    plt.xlabel('accumulated accounts')
    plt.ylabel('commitment time/ns')
    plt.title('Ethereum StateDB insertion for 10000W accounts')
    plt.show()
    plt.savefig('file/output_2.png')


if __name__ == '__main__':
    csv1()
    csv2()
