package bet

import (
	"bytes"
	"encoding/json"
	"html/template"
)

type servedBet struct {
	Bet
	FormattedDate string
}

func ServeLastBets() (string, error) {
	bets, err := getLastBets("bets")
	totalPages := (len(bets) / 15) + 1
	if err != nil {
		return "", err
	}
	var s []servedBet
	for _, b := range bets {
		s = append(s, servedBet{
			Bet:           b,
			FormattedDate: b.Posted_at.Format("2006-01-02 15:04:05"),
		})
	}
	htmlTable := generateTable(s, 1, totalPages, 15)
	return htmlTable, nil
}

func generateTable(bets []servedBet, currentPage int, totalPages int, itemsperPage int) string {
	const tableTemplate = `
        <!DOCTYPE html>
        <html>
        <head>
            <title>LEGROUP</title>
            <link rel="stylesheet" type="text/css" href="/static/styles.css">
        </head>
        <body>
        <div id="table-container">
            <table>
                <tr>
                    <th>team</th>
                    <th>pick</th>
                    <th>size</th>
                    <th>odds</th>
                    <th>result</th>
                    <th>date</th>
                </tr>
                {{range .Bets}}
                <tr class="bet-row">
                    <td>{{.Team}}</td>
                    <td>{{.Prediction}}</td>
                    <td>{{.Size}}</td>
                    <td>{{.Odds}}</td>
                    <td class="result">{{.Result}}</td>
                    <td>{{.FormattedDate}}</td>
                </tr>
                {{end}}
            </table>
        </div>
        <div id="pagination-container">
            <ul id="pagination"></ul>
        </div>
        <script>
            // Number of items per page
            var itemsPerPage = {{.ItemsPerPage}};

            // Your data, replace with your actual data
            var data = {{.Bets | jsonify}};

            var currentPage = {{.CurrentPage}};
            var totalPages = {{.TotalPages}};

            function displayTablePage(page) {
                var startIndex = (page - 1) * itemsPerPage;
                var endIndex = page * itemsPerPage;
                var tableData = data.slice(startIndex, endIndex);

                var tableBody = "<tr><th>team</th><th>pick</th><th>size</th><th>odds</th><th>result</th><th>date</th></tr>";

                for (var i = 0; i < tableData.length; i++) {
                    var resultClass = tableData[i].Result === "won" ? "won" : "lost";
                    tableBody += "<tr class='bet-row " + resultClass + "'>";
                    tableBody += "<td>" + tableData[i].Team + "</td>";
                    tableBody += "<td>" + tableData[i].Prediction + "</td>";
                    tableBody += "<td>" + tableData[i].Size + "</td>";
                    tableBody += "<td>" + tableData[i].Odds + "</td>";
                    tableBody += "<td class='result'>" + tableData[i].Result + "</td>";
                    tableBody += "<td>" + tableData[i].FormattedDate + "</td>";
                    tableBody += "</tr>";
                }

                document.querySelector("#table-container table").innerHTML = tableBody;

                var rows = document.querySelectorAll(".bet-row");
                rows.forEach(row => {
                    const result = row.querySelector(".result").textContent;
                    if (result === "won") {
                        row.classList.add("won");
                    } else {
                        row.classList.add("lost");
                    }
                });
            }

            function displayPagination() {
                var pagination = document.getElementById("pagination");
                pagination.innerHTML = "";

                for (var i = 1; i <= totalPages; i++) {
                    var li = document.createElement("li");
                    li.innerHTML = i;
                    li.onclick = function () {
                        currentPage = parseInt(this.innerHTML);
                        displayTablePage(currentPage);
                        highlightCurrentPage();
                    };
                    pagination.appendChild(li);
                }

                highlightCurrentPage();
            }

            function highlightCurrentPage() {
                var pages = document.getElementById("pagination").getElementsByTagName("li");
                for (var i = 0; i < pages.length; i++) {
                    pages[i].classList.remove("active");
                }
                pages[currentPage - 1].classList.add("active");
            }

            displayTablePage(currentPage);
            displayPagination();
        </script>

        </body>
        </html>
    `

	t := template.Must(template.New("table").Funcs(template.FuncMap{"jsonify": jsonify}).Parse(tableTemplate))

	var buf bytes.Buffer
	t.Execute(&buf, struct {
		Bets         []servedBet
		CurrentPage  int
		TotalPages   int
		ItemsPerPage int
	}{bets, currentPage, totalPages, itemsperPage})

	return buf.String()
}

func jsonify(data interface{}) template.JS {
	jsonData, _ := json.Marshal(data)
	return template.JS(jsonData)
}
