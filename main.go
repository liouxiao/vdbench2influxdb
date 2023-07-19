package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	host := "127.0.0.1"
	port := "8086"
	db := "G200"

	url := "http://" + host + ":" + port + "/write?db=" + db + "&precision=ms"

	tag := "n/a"
	if len(os.Args) > 2 {
		tag = os.Args[2]
	}

	if len(os.Args) > 3 {
		url = os.Args[3]
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	count := 0
	data := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		iteration := strings.Fields(line)
		if len(iteration) != 34 {
			continue
		}

		// tod timestamp Run Interval reqrate rate MB/sec bytes/io
		// read% resp read_resp write_resp resp_max read_max write_max
		// resp_std read_std write_std xfersize threads rdpct rhpct
		// whpct seekpct lunsize version compratio dedupratio
		// queue_depth cpu_used cpu_user cpu_kernel cpu_wait cpu_idle
		tod, timestamp, run, _, reqrate, rate, mbsec, bytesio,
			readpct, resp, readResp, writeResp, respMax, readMax, writeMax,
			respStd, readStd, writeStd, xfersize, threads, rdpct, rhpct,
			whpct, seekpct, lunsize, version, compratio, dedupratio,
			queueDepth, cpuUsed, cpuUser, cpuKernel, cpuWait, cpuIdle :=
			iteration[0], iteration[1], iteration[2], iteration[3], iteration[4], iteration[5], iteration[6], iteration[7],
			iteration[8], iteration[9], iteration[10], iteration[11], iteration[12], iteration[13], iteration[14],
			iteration[15], iteration[16], iteration[17], iteration[18], iteration[19], iteration[20], iteration[21],
			iteration[22], iteration[23], iteration[24], iteration[25], iteration[26], iteration[27],
			iteration[28], iteration[29], iteration[30], iteration[31], iteration[32], iteration[33]

		if timestamp == "timestamp" {
			continue
		}

		// timestamp = strings.ReplaceAll(timestamp, "-", " ")
		ms := strings.SplitN(tod, ".", 2)[1]
		parsedTime, _ := time.Parse("01/02/2006-15:04:05-MST", timestamp) // "2006-01-02 15:04:05", timestamp)
		timestamp = fmt.Sprintf("%d%s", parsedTime.UnixNano()/int64(time.Second), ms)

		data += fmt.Sprintf("reqrate,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(reqrate), timestamp)
		data += fmt.Sprintf("rate,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(rate), timestamp)
		data += fmt.Sprintf("mbsec,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(mbsec), timestamp)
		data += fmt.Sprintf("bytesio,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(bytesio), timestamp)
		data += fmt.Sprintf("readpct,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(readpct), timestamp)
		data += fmt.Sprintf("resp,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(resp), timestamp)
		data += fmt.Sprintf("read_resp,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(readResp), timestamp)
		data += fmt.Sprintf("write_resp,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(writeResp), timestamp)
		data += fmt.Sprintf("resp_max,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(respMax), timestamp)
		data += fmt.Sprintf("read_max,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(readMax), timestamp)
		data += fmt.Sprintf("write_max,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(writeMax), timestamp)
		data += fmt.Sprintf("resp_std,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(respStd), timestamp)
		data += fmt.Sprintf("read_std,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(readStd), timestamp)
		data += fmt.Sprintf("write_std,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(writeStd), timestamp)
		data += fmt.Sprintf("xfersize,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(xfersize), timestamp)
		data += fmt.Sprintf("threads,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(threads), timestamp)
		data += fmt.Sprintf("rdpct,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(rdpct), timestamp)
		data += fmt.Sprintf("rhpct,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(rhpct), timestamp)
		data += fmt.Sprintf("whpct,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(whpct), timestamp)
		data += fmt.Sprintf("seekpct,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(seekpct), timestamp)
		data += fmt.Sprintf("lunsize,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(lunsize), timestamp)
		data += fmt.Sprintf("version,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(version), timestamp)
		data += fmt.Sprintf("compratio,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(compratio), timestamp)
		data += fmt.Sprintf("dedupratio,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(dedupratio), timestamp)
		data += fmt.Sprintf("queue_depth,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(queueDepth), timestamp)
		data += fmt.Sprintf("cpu_used,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(cpuUsed), timestamp)
		data += fmt.Sprintf("cpu_user,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(cpuUser), timestamp)
		data += fmt.Sprintf("cpu_kernel,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(cpuKernel), timestamp)
		data += fmt.Sprintf("cpu_wait,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(cpuWait), timestamp)
		data += fmt.Sprintf("cpu_idle,tag=%s,run=%s value=%.2f %s\n", tag, run, parseFloat(cpuIdle), timestamp)

		count = count + 1

		// write to influxdb every 5 mins
		if count == 300 {
			writeInflux(url, data)
			data = ""
			count = 0
		}
	}

	if data != "" {
		writeInflux(url, data)
		data = ""
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

}

func parseFloat(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}

func writeInflux(url, data string) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println("Error sending the request:", err)
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return
	}

	fmt.Print(".")
}
