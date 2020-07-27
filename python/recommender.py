import pandas as pd
import numpy as np
import re
import json
import os
from time import sleep
from google_drive_downloader import GoogleDriveDownloader as gdd

if os.path.exists("./moviemat.csv"):
	print("moviemat.csv -> ok")
else:
	print("moviemat.csv -> downloading")
	gdd.download_file_from_google_drive(file_id='1t-ogbQRTTGfjbdNMwoGsnFzB5rQMqUkU', dest_path='./moviemat.zip', unzip=True)

moviemat = pd.read_csv("moviemat.csv", dtype=np.float16)
movies = moviemat.columns.to_list()
movies_lower = [movie.lower() for movie in movies]
movies_lower = [re.sub(r'[^0-9^a-z^\s]', '', movie) for movie in movies_lower]

movie_dict = {
	"movies": movies
}

json_string = json.dumps(movie_dict)

with open('/rest/titles.json', mode='w', encoding="utf-8") as outfile:
	outfile.write(json_string)

del movie_dict
print("Titles saved!")

theNumber = 0
user_movie = ""
json_string = ""

# Ensure not to read the file at start (value 0 in the read file)
with open("../rest/read", mode='w') as f:
	f.write(str(theNumber))

while True:

	with open("../rest/read", mode='r') as f:
		theNumber = int(f.readline())

	if theNumber == 1:
		with open('/rest/usermovie.json', mode='r', encoding="utf-8") as f:
			user_movie = json.load(f)['movie']
			user_movie = user_movie.lower()
			user_movie = re.sub(r'[^0-9^a-z^\s]', '', user_movie)

		r = re.compile(user_movie + "*")

		try:
			movie_name = list(filter(r.match, movies_lower))[0]
			movie_idx = movies_lower.index(movie_name)
			movie_name = movies[movie_idx]
			print(f"\n\rThe selected movie is: {movie_name}\n\r")

			# Recommender
			sub_ds = moviemat[movie_name]
			similars = moviemat.corrwith(sub_ds)
			similars = similars.astype(np.float16)
			df01 = pd.DataFrame(data=similars, columns=["Pearson"])
			df01 = df01[df01["Pearson"] > 0]
			df01.sort_values("Pearson", ascending=False, inplace=True)
			recommended_movies = df01.index.to_list()[1:4]
			df01 = pd.DataFrame(None)

			print("We recommend you these movies:")

			for i in range(0, len(recommended_movies)):
				print(f"{i+1}- {recommended_movies[i]}")
			print("\n\r")

			result = {
				"movie": movie_name,
				"recommend": recommended_movies
			}
			json_string = json.dumps(result)

		except:
			print(
				"\n\rError: The movie is not in our database. Please try another one.\n\r")
			result = {
				"movie": "The movie is not in our database. Please try another one."
			}
			json_string = json.dumps(result)

		# save number and JSON
		theNumber = 0
		with open("../rest/read", mode='w') as f:
			f.write(str(theNumber))

		# JSON
		with open('../rest/rcmd_movies.json', mode='w', encoding="utf-8") as outfile:
			outfile.write(json_string)

	sleep(0.10)
