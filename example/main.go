package main

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/l1va/gofins/fins"
)

func main() {

	clientAddr := fins.NewAddress("192.168.22.5", 9600, 0, 5, 0) // PC's IP
	plcAddr := fins.NewAddress("192.168.22.3", 9600, 0, 3, 0) //PLC's IP
	//clientAddr := fins.NewAddress("127.0.0.1", 9600, 0, 34, 0)
	//plcAddr := fins.NewAddress("127.0.0.1", 9600, 0, 0, 0)

	/******PLCSimulator******
	s, e := fins.NewPLCSimulator(plcAddr)
	if e != nil {
		panic(e)
	}
	defer s.Close()
	*/

	c, err := fins.NewClient(clientAddr, plcAddr) //创建一个通讯客户端
	if err != nil {
		panic(err)
	}
	defer c.Close() //主函数退出时执行
	
	//读取PLC Clock*****测试成功*****
	time, _ := c.ReadClock()
	fmt.Println(time)
	//2020-11-10 11:19:10 +0800 CST

	//读写word****测试成功*****
	c.WriteWords(fins.MemoryAreaWRWord, 100, []uint16{23, 24})
	z, err := c.ReadWords(fins.MemoryAreaWRWord, 100, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(z)

	//读写Bit****测试成功*****
	c.WriteBits(fins.MemoryAreaWRBit, 0, 0, []bool{true, false, true, false}) //参数：memory类型、word起始地址、偏移量、写入的值
	//c.SetBit(fins.MemoryAreaWRBit, 0, 0) //bit置1，参数：memory类型、word起始地址、偏移量
	//c.ResetBit(fins.MemoryAreaWRBit, 0, 0) //bit置0，参数：memory类型、word起始地址、偏移量
	//c.ToggleBit(fins.MemoryAreaWRBit, 0, 0) //bit值翻转，参数：memory类型、word起始地址、偏移量
	bit, _ := c.ReadBits(fins.MemoryAreaWRBit, 0, 0, 4) //参数：memory类型、word起始地址、偏移量、读取数量
	fmt.Println(bit)

	//读取输入CIO 0.0*****测试成功*****
	bit, _ = c.ReadBits(fins.MemoryAreaCIOBit, 0, 0, 4) //参数：memory类型、word起始地址、偏移量、读取数量
	fmt.Println(bit)

	//读取输出CIO 100.0*****测试成功*****
	bit, _ = c.ReadBits(fins.MemoryAreaCIOBit, 100, 0, 4) //参数：memory类型、word起始地址、偏移量、读取数量
	fmt.Println(bit)

	c.ToggleBit(fins.MemoryAreaCIOBit, 0, 0) //bit值翻转，参数：memory类型、word起始地址、偏移量
	//读取输出CIO 100.0*****测试成功*****
	bit, _ = c.ReadBits(fins.MemoryAreaCIOBit, 100, 0, 4) //参数：memory类型、word起始地址、偏移量、读取数量
	fmt.Println(bit)

	//读取输出CIO 100.0 开始的word*****测试成功*****
	z, err = c.ReadWords(fins.MemoryAreaCIOWord, 100, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(z)

	
	/*
				//测试读写word
				c.WriteWords(fins.MemoryAreaDMWord, 2000, []uint16{z[0] + 1, z[1] - 1})

				z, err = c.ReadWords(fins.MemoryAreaDMWord, 2000, 50)
				if err != nil {
					panic(err)
				}
				fmt.Println(z)
				// output: [1 65535 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

				//测试读写float
				buf := make([]byte, 8, 8)
				binary.LittleEndian.PutUint64(buf[:], math.Float64bits(15.6))
				err = c.WriteBytes(fins.MemoryAreaDMWord, 10, buf)
				if err != nil {
					panic(err)
				}

				b, err := c.ReadBytes(fins.MemoryAreaDMWord, 10, 4)
				if err != nil {
					panic(err)
				}
				floatRes := math.Float64frombits(binary.LittleEndian.Uint64(b))
				fmt.Println("Float result:", floatRes)
				// output: Float result: 15.6

				// 测试读写string
				err = c.WriteString(fins.MemoryAreaDMWord, 10000, "teststring")
				if err != nil {
					panic(err)
				}

				str, _ = c.ReadString(fins.MemoryAreaDMWord, 10000, 5)
				fmt.Println(str, len(str))
				// output: teststring 10

		                //测试读写bit
				bit, _ := c.ReadBits(fins.MemoryAreaDMWord, 10473, 2, 1)
				fmt.Println(bit)

				c.WriteBits(fins.MemoryAreaDMBit, 24002, 0, []bool{false, false, false, true})
				c.SetBit(fins.MemoryAreaDMBit, 24003, 1)
				c.ResetBit(fins.MemoryAreaDMBit, 24003, 1)
				c.ToggleBit(fins.MemoryAreaDMBit, 24003, 1)

				cron := cron.New()
				s := rasc.NewShelter()
				cron.AddFunc("*5 * * * * *", func() {
					t, _ := c.ReadClock()
					fmt.Printf("Setting PLC time to: %s\n", t.Format(time.RFC3339))
					c.WriteString(fins.MemoryAreaDMWord, 10000, 10, t.Format(time.RFC3339))
				})
				cron.Start()
	*/
}
