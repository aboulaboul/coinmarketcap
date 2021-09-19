# coinmarketcap : utility to get cryptos quote (EUR or USD) & market capitalization from coinmarketcap.com API

# install
<p>
git clone git@github.com:aboulaboul/coinmarketcap.git
</p>

# build
<p>
cd coinmarketcap
go build
</p>

# run minimal option under linux (con't : you need a personnal API key from coinmarketcap)
<p>
./coinmarketcap -api-key=xxxxxxxxxxxxxxx
</p>

# run minimal option with export to file  under linux
<p>
./coinmarketcap -api-key=xxxxxxxxxxxxxxx > myfile.csv
</p>

# help under linux
<p>
./coinmarketcap --help
</p>