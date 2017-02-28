package kaa

const gameString1 = `{"you":"0623b12a-411b-4674-a115-591063ef92d3","width":10,"turn":124,"snakes":[{"taunt":"battlesnake-go!","name":"7eef72e9-72fc-4c27-a387-898384639f46 (10x10)","id":"0623b12a-411b-4674-a115-591063ef92d3","health_points":96,"coords":[[9,1],[9,0],[8,0],[8,1],[8,2],[7,2],[7,3],[7,4],[7,5],[7,6],[6,6],[6,7],[5,7],[4,7],[3,7],[2,7],[1,7],[1,8],[0,8],[0,7],[0,6],[1,6],[2,6],[3,6],[4,6],[5,6],[5,5],[5,4],[5,3],[5,2],[6,2],[6,1]]}],"height":10,"game_id":"7eef72e9-72fc-4c27-a387-898384639f46","food":[[0,0],[1,3],[4,0]],"dead_snakes":[]}`

const gameString2 = `{"you":"82557bbc-5ff2-4e51-8133-f6875d4f8d71","width":10,"turn":233,"snakes":[{"taunt":"battlesnake-go!","name":"7eef72e9-72fc-4c27-a387-898384639f46 (10x10)","id":"82557bbc-5ff2-4e51-8133-f6875d4f8d71","health_points":100,"coords":[[1,3],[0,3],[0,4],[0,5],[0,6],[0,7],[1,7],[2,7],[3,7],[3,8],[3,9],[4,9],[4,8],[4,7],[4,6],[4,5],[5,5],[5,4],[6,4],[7,4],[7,3],[6,3],[5,3],[4,3],[4,4],[3,4],[3,3],[3,2],[4,2],[5,2],[5,1],[4,1],[3,1],[2,1],[2,0],[3,0],[4,0],[5,0],[6,0],[7,0],[8,0],[9,0],[9,1],[9,2],[9,2]]}],"height":10,"game_id":"7eef72e9-72fc-4c27-a387-898384639f46","food":[[6,2],[7,5],[2,3]],"dead_snakes":[]}`

const gameString3 = `{"you":"82557bbc-5ff2-4e51-8133-f6875d4f8d71","width":10,"turn":223,"snakes":[{"taunt":"battlesnake-go!","name":"7eef72e9-72fc-4c27-a387-898384639f46 (10x10)","id":"82557bbc-5ff2-4e51-8133-f6875d4f8d71","health_points":100,"coords":[[3,9],[4,9],[4,8],[4,7],[4,6],[4,5],[5,5],[5,4],[6,4],[7,4],[7,3],[6,3],[5,3],[4,3],[4,4],[3,4],[3,3],[3,2],[4,2],[5,2],[5,1],[4,1],[3,1],[2,1],[2,0],[3,0],[4,0],[5,0],[6,0],[7,0],[8,0],[9,0],[9,1],[9,2],[9,3],[9,4],[9,5],[9,6],[9,7],[9,8],[8,8],[8,8]]}],"height":10,"game_id":"7eef72e9-72fc-4c27-a387-898384639f46","food":[[0,5],[0,7],[0,6]],"dead_snakes":[]}`

const gameString4 = `{"you":"0086b0de-4d23-4189-9305-a7308240edb4","width":20,"turn":176,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"3a1cf2b6-ab7f-4870-b672-cb60d7ab4e67","health_points":90,"coords":[[1,15],[2,15],[3,15],[4,15],[4,14],[4,13],[4,12],[4,11],[4,10],[4,9],[3,9],[3,10],[3,11],[2,11],[1,11],[1,10],[1,9],[0,9],[0,8],[1,8],[2,8],[2,7],[3,7],[3,6],[3,5],[4,5],[5,5],[6,5],[7,5],[8,5],[9,5],[9,6],[9,7],[10,7]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"0086b0de-4d23-4189-9305-a7308240edb4","health_points":94,"coords":[[14,16],[13,16],[12,16],[11,16],[10,16],[10,15],[10,14],[11,14],[11,13],[10,13],[9,13],[8,13],[8,14],[8,15],[8,16],[8,17],[8,18],[7,18],[6,18],[6,19],[7,19],[8,19],[9,19],[10,19],[10,18],[11,18],[12,18],[13,18],[14,18],[15,18],[16,18],[17,18],[17,17],[17,16],[17,15],[16,15],[15,15],[15,14],[14,14],[13,14]]}],"height":20,"game_id":"1dd0a217-baef-46ca-bd17-d48a38d54436","food":[[0,6],[16,16],[15,19],[18,7],[8,4],[10,2],[0,15],[3,1],[13,2],[14,11]],"dead_snakes":[]}`

const gameString5 = `{"you":"3de4f206-1538-4bfc-ad49-4cb17fc61bb5","width":10,"turn":8,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"3de4f206-1538-4bfc-ad49-4cb17fc61bb5","health_points":92,"coords":[[5,8],[6,8],[7,8]]}],"height":10,"game_id":"8d58c168-aa8c-4395-90dc-79e36b32bf1e","food":[[2,5],[5,7],[1,8],[8,7],[6,2],[2,8],[7,4],[6,3],[4,2],[6,4]],"dead_snakes":[]}`

const gameString6 = `{"you":"d6f6ff70-fc64-4277-a4a8-b9d0106a87ce","width":20,"turn":376,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"d6f6ff70-fc64-4277-a4a8-b9d0106a87ce","health_points":79,"coords":[[0,13],[0,14],[1,14],[1,15],[0,15],[0,16],[0,17],[0,18],[0,19],[1,19],[1,18],[1,17],[1,16],[2,16],[2,15],[3,15],[3,14],[4,14],[5,14],[6,14],[7,14],[7,15],[6,15],[5,15],[5,16],[5,17],[5,18],[5,19],[6,19],[6,18],[7,18],[8,18],[9,18],[10,18],[10,17],[9,17],[8,17],[8,16],[8,15],[8,14],[8,13],[8,12],[8,11],[7,11],[6,11],[5,11],[4,11],[3,11],[2,11],[1,11],[0,11],[0,10],[0,9],[0,8],[0,7],[0,6],[0,5],[0,4],[0,3],[0,2],[0,1],[1,1],[2,1]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[14,3],[9,19],[17,14],[12,5],[11,6],[12,8]],"dead_snakes":[]}`

const gameString7 = `{"you":"c38c3d40-681c-4f37-8060-bc954e5d1005","width":20,"turn":233,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"c38c3d40-681c-4f37-8060-bc954e5d1005","health_points":96,"coords":[[11,10],[11,9],[11,8],[11,7],[11,6],[12,6],[13,6],[13,5],[13,4],[13,3],[13,2],[13,1],[13,0],[12,0],[11,0],[10,0],[9,0],[9,1],[9,2],[9,3],[9,4],[9,5],[9,6],[9,7],[9,8],[9,9],[9,10],[9,11],[10,11],[11,11],[11,12],[11,13],[11,14],[11,15],[12,15],[13,15]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[6,19],[0,18],[2,13],[0,17],[7,19],[2,1]],"dead_snakes":[]}`

const gameString8 = `{"you":"c08ce082-1ea3-47b1-b347-fd7faa42bdd8","width":20,"turn":606,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"c08ce082-1ea3-47b1-b347-fd7faa42bdd8","health_points":96,"coords":[[16,1],[17,1],[18,1],[19,1],[19,0],[18,0],[17,0],[16,0],[15,0],[14,0],[13,0],[12,0],[11,0],[11,1],[11,2],[11,3],[11,4],[11,5],[11,6],[12,6],[13,6],[13,7],[13,8],[12,8],[11,8],[10,8],[9,8],[8,8],[7,8],[6,8],[6,9],[6,10],[7,10],[7,9],[8,9],[9,9],[10,9],[11,9],[12,9],[13,9],[14,9],[15,9],[16,9],[17,9],[18,9],[18,8],[19,8],[19,9],[19,10],[19,11],[19,12],[18,12],[17,12],[16,12],[15,12],[14,12],[14,11],[14,10],[13,10],[12,10],[12,11],[11,11],[10,11],[9,11],[8,11],[7,11],[7,12],[7,13],[7,14],[8,14],[8,15],[8,16],[7,16],[6,16],[5,16],[4,16],[3,16],[2,16],[2,15],[2,14],[3,14]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[2,11],[2,13],[14,13],[2,10],[7,18],[15,19]],"dead_snakes":[]}
`
const gameString9 = `{"you":"f30c5778-b50b-45a3-8d17-05117ebd35a9","width":20,"turn":125,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"40bae8a0-5d41-4dd4-ae3a-3496eb9a6603","health_points":92,"coords":[[9,2],[10,2],[11,2],[12,2],[12,3],[12,4],[13,4],[14,4],[14,5],[13,5],[12,5],[12,6],[11,6],[11,5],[11,4],[11,3],[10,3],[9,3],[8,3],[7,3],[6,3]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"f30c5778-b50b-45a3-8d17-05117ebd35a9","health_points":89,"coords":[[13,2],[13,3],[14,3],[15,3],[15,4],[15,5],[15,6],[15,7],[14,7],[13,7],[12,7],[11,7],[11,8],[11,9],[11,10],[11,11],[11,12],[11,13],[11,14],[11,15],[11,16],[11,17]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"7a996f7c-2a92-4327-a1d3-464931ae4aa2","health_points":89,"coords":[[13,16],[13,15],[13,14],[13,13],[13,12],[13, 11],[13,10],[13,9],[14,9],[14,8],[15,8],[16,8],[16,9],[16,10],[15,10],[14,10],[14,11],[14,12],[14,13]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[4,18],[5,4],[1,11],[3,16],[0,11],[8,0]],"dead_snakes":[]}
`
const gameString10 = `{"you":"301ac963-6fd0-478b-b1d1-7d9a227b07d4","width":20,"turn":199,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"0e0d9e0d-5ff7-4cc8-8052-cf49d92bc347","health_points":97,"coords":[[1,6],[1,7],[1,8],[1,9],[0,9],[0,10],[0,11],[0,12],[1,12],[2,12],[3,12],[4,12],[5,12],[6,12],[7,12],[8,12],[9,12],[9,11],[9,10],[9,9],[10,9],[11,9],[11,8],[11,7],[11,6],[10,6],[9,6],[8,6],[8,5],[8,4],[8,3]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"e3e306d5-f5c0-4dc0-9b5e-6896f8ebcd9a","health_points":96,"coords":[[5,19],[4,19],[4,18],[3,18],[3,17],[4,17],[4,16],[5,16],[6,16],[7,16],[8,16],[9,16],[9,15],[9,14],[9,13],[10,13],[10,12],[10,11],[10,10],[11,10]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"301ac963-6fd0-478b-b1d1-7d9a227b07d4","health_points":100,"coords":[[7,18],[7,17],[8,17],[9,17],[10,17],[10,16],[11,16],[12,16],[13,16],[14,16],[15,16],[16,16],[17,16],[18,16],[18,15],[17,15],[16,15],[16,14],[16,13],[16,12],[15,12],[14,12],[14,11],[14,11]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[15,7],[12,18],[5,6],[18,3],[5,1],[16,2]],"dead_snakes":[]}
`

const gameString11 = `{"you":"192c3b82-5424-4092-9284-bbc567e3d25f","width":20,"turn":1145,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"192c3b82-5424-4092-9284-bbc567e3d25f","health_points":99,"coords":[[8,19],[9,19],[9,18],[9,17],[9,16],[9,15],[8,15],[8,14],[9,14],[10,14],[11,14],[12,14],[12,15],[13,15],[14,15],[15,15],[15,14],[15,13],[15,12],[15,11],[16,11],[17,11],[18,11],[19,11],[19,10],[19,9],[19,8],[18,8],[17,8],[16,8],[15,8],[15,7],[15,6],[15,5],[16,5],[17,5],[17,4],[16,4],[15,4],[14,4],[14,3],[14,2],[13,2],[12,2],[12,1],[11,1],[10,1],[9,1],[8,1],[7,1],[6,1],[5,1],[5,0],[4,0],[4,1],[4,2],[3,2],[3,3],[3,4],[2,4],[1,4],[1,5],[1,6],[1,7],[1,8],[1,9],[1,10],[1,11],[2,11],[3,11],[4,11],[4,10],[5,10],[5,9],[5,8],[6,8],[6,9],[6,10],[6,11],[5,11],[5,12],[5,13],[5,14],[5,15],[4,15],[4,16],[3,16],[2,16],[1,16],[1,17],[1,18],[1,19],[2,19],[2,18],[2,17],[3,17],[3,18],[3,19],[4,19],[5,19],[5,18],[5,17],[5,16],[6,16],[6,15],[6,14],[6,13],[6,12],[7,12],[8,12],[9,12],[10,12],[11,12],[12,12],[13,12],[13,13],[13,14],[14,14],[14,13],[14,12],[14,11],[14,10],[15,10],[16,10],[17,10],[18,10],[18,9],[17,9],[16,9],[15,9],[14,9],[14,8]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[5,3],[13,1],[4,7],[15,0],[0,2],[6,19]],"dead_snakes":[{"taunt":"Obviously you're not a golfer.","name":"Rouser","id":"a174bb49-e6d2-4841-9650-b71f96cf7127","health_points":91,"coords":[[7,10],[7,11],[7,10],[7,9],[6,9],[5,9],[4,9],[4,10],[5,10],[6,10],[6,11],[5,11],[4,11],[4,12],[4,13],[3,13],[2,13],[1,13],[0,13],[0,12],[1,12],[1,11],[0,11],[0,10],[0,9],[0,8],[1,8],[2,8],[3,8],[4,8],[5,8],[6,8],[7,8],[8,8],[8,9],[8,10],[8,11],[8,12],[8,13],[8,14],[9,14],[10,14],[11,14],[12,14],[12,13],[12,12],[12,11],[13,11],[13,10],[14,10],[15,10],[16,10],[16,11],[15,11],[15,12],[14,12]]},{"taunt":"Obviously you're not a golfer.","name":"Rouser","id":"70b2b4eb-5b4e-4dc3-949a-d6716591d0cc","health_points":99,"coords":[[12,2],[12,3],[11,3],[11,4],[12,4],[12,5],[13,5],[13,6],[13,7],[13,8],[12,8],[11,8],[10,8],[9,8],[8,8],[8,7],[8,6],[8,5],[8,4],[8,3],[8,2],[9,2],[10,2],[11,2],[12,2],[13,2],[14,2],[14,1],[15,1],[16,1],[16,0],[17,0],[18,0],[19,0],[19,1],[19,2],[19,3],[19,4],[19,5],[18,5]]}]}`

const gameString12 = `{"you":"67c6ff3c-682b-43c2-b055-2449419eb753","width":20,"turn":13,"snakes":[{"taunt":"Obviously you're not a golfer.","name":"Rouser","id":"10240f8a-daeb-4136-a811-3c0799752b0c","health_points":98,"coords":[[13,11],[13,12],[12,12],[12,11],[11,11],[10,11]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"67c6ff3c-682b-43c2-b055-2449419eb753","health_points":92,"coords":[[7,11],[7,10],[6,10],[5,10]]},{"taunt":"Keen and green: (add (add (add (add var0 (min var2 var0)) (add (add (neg var1) (min var2 var0)) (add (add 2.50242369095 (min var2 var0)) (min (add (add (neg var1) (min var2 var0)) (min var2 var0)) var0)))) (min (add (add (add (neg var1) (min var2 var0)) (add (add (neg var1) (min var2 var0)) (add (add 2.50242369095 (min var2 var0)) (min (add (add (neg var1) (min var2 var0)) (min var2 var0)) var0)))) var0) var0)) (add (add (add (neg var1) (min var2 var0)) (add (add (neg var1) (min var2 var0)) (add (add 2.50242369095 (min var2 var0)) (min (add (add (neg var1) (min var2 var0)) (min var2 var0)) var0)))) var0))","name":"NBK Trainee","id":"e2c601a0-8efd-45a2-b850-d0a6b168b043","health_points":91,"coords":[[16,19],[15,19],[14,19],[13,19]]},{"taunt":"I'm a snake, not a doctor","name":"Genetisnake NBK","id":"90095ff2-affc-4fed-b797-6eeecb055d20","health_points":100,"coords":[[2,4],[3,4],[3,5],[3,6],[3,6]]},{"taunt":"hungry hungry snake: (neg var1)","name":"NBK Greedy","id":"ee5f46e9-c476-4541-9efb-c9455d12abfb","health_points":100,"coords":[[12,8],[12,7],[12,6],[12,5],[12,4],[12,4]]}],"height":20,"game_id":"01d41d8a-ae1c-4334-9c2c-0e0ee9dcbb76","food":[[7,13],[2,2],[9,9],[16,6],[13,10],[17,1]],"dead_snakes":[]}`

const gameString13 = `{"you":"e62875a9-2cac-4081-99b9-d0f1665503d8","width":20,"turn":6,"snakes":[{"taunt":"Obviously you're not a golfer.","name":"Rouser","id":"69fc60c4-d343-4590-a6ca-de36450c47c9","health_points":94,"coords":[[12,8],[13,8],[14,8]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"e62875a9-2cac-4081-99b9-d0f1665503d8","health_points":94,"coords":[[8,7],[8,6],[8,5]]},{"taunt":"Keen and green: (add (add (add (add var0 (min var2 var0)) (add (add (neg var1) (min var2 var0)) (add (add 2.50242369095 (min var2 var0)) (min (add (add (neg var1) (min var2 var0)) (min var2 var0)) var0)))) (min (add (add (add (neg var1) (min var2 var0)) (add (add (neg var1) (min var2 var0)) (add (add 2.50242369095 (min var2 var0)) (min (add (add (neg var1) (min var2 var0)) (min var2 var0)) var0)))) var0) var0)) (add (add (add (neg var1) (min var2 var0)) (add (add (neg var1) (min var2 var0)) (add (add 2.50242369095 (min var2 var0)) (min (add (add (neg var1) (min var2 var0)) (min var2 var0)) var0)))) var0))","name":"NBK Trainee","id":"5b0374c1-8a9f-4fd6-9954-6c75af1f3b82","health_points":94,"coords":[[10,11],[10,12],[10,13]]},{"taunt":"I'm a snake, not a doctor","name":"Genetisnake NBK","id":"55c079e6-6212-42f6-86aa-962056fce571","health_points":94,"coords":[[19,5],[19,6],[19,7]]},{"taunt":"hungry hungry snake: (neg var1)","name":"NBK Greedy","id":"2445620a-fe3f-49ae-9f1f-6330cccfdee4","health_points":100,"coords":[[14,17],[15,17],[15,16],[16,16],[16,16]]}],"height":20,"game_id":"01d41d8a-ae1c-4334-9c2c-0e0ee9dcbb76","food":[[10,2],[5,8],[16,4],[18,4],[14,4],[10,8]],"dead_snakes":[]}`

const gameString14 = `{"you":"6c4bdcf4-b9fa-4f23-92c1-42df7bc3ba00","width":20,"turn":143,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"6c4bdcf4-b9fa-4f23-92c1-42df7bc3ba00","health_points":99,"coords":[[16,3],[16,4],[15,4],[14,4],[13,4],[13,5],[13,6],[12,6],[12,7],[12,8],[11,8],[11,9],[12,9],[12,10],[12,11],[12,12],[12,13],[12,14],[11,14],[10,14],[10,15],[11,15],[11,16]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"de50b5ba-c887-458c-ad9c-0c61e0a968e0","health_points":100,"coords":[[9,14],[8,14],[7,14],[6,14],[5,14],[4,14],[3,14],[2,14],[2,15],[2,16],[1,16],[0,16],[0,15],[0,14],[0,13],[1,13],[2,13],[2,13]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"d6f783bd-401d-4ccc-918f-9ed2536f2810","health_points":93,"coords":[[15,2],[16,2],[16,1],[15,1],[14,1],[13,1],[12,1],[11,1],[11,0],[12,0],[13,0],[14,0],[15,0],[16,0],[17,0],[17,1],[17,2],[17,3],[17,4],[17,5],[16,5],[15,5],[14,5]]}],"height":20,"game_id":"01d41d8a-ae1c-4334-9c2c-0e0ee9dcbb76","food":[[5,9],[10,12],[2,12],[3,12],[17,15],[18,13]],"dead_snakes":[]}`

const gameString15 = `{"you":"97dacadc-e9cd-421d-b316-0b36159d33ef","width":20,"turn":629,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"97dacadc-e9cd-421d-b316-0b36159d33ef","health_points":97,"coords":[[6,19],[6,18],[6,17],[6,16],[7,16],[8,16],[8,15],[9,15],[9,14],[10,14],[11,14],[11,13],[10,13],[9,13],[8,13],[8,14],[7,14],[6,14],[5,14],[4,14],[3,14],[2,14],[2,13],[2,12],[2,11],[1,11],[1,10],[1,9],[1,8],[2,8],[2,7],[2,6],[3,6],[3,7],[3,8],[3,9],[2,9],[2,10],[3,10],[4,10],[5,10],[6,10],[7,10],[7,9],[8,9],[8,8],[8,7],[9,7],[10,7],[10,6],[11,6],[12,6],[13,6],[14,6],[15,6],[15,7],[15,8],[16,8],[17,8],[17,9],[17,10],[17,11],[17,12],[17,13],[17,14],[16,14],[15,14],[14,14],[14,15],[14,16],[14,17],[14,18],[15,18],[16,18],[17,18],[17,19],[16,19],[15,19],[14,19],[13,19],[13,18],[13,17],[13,16],[13,15],[13,14],[13,13],[14,13]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"959fd111-95d7-40b6-82ec-ed70d8939552","health_points":99,"coords":[[8,3],[8,4],[7,4],[6,4],[5,4],[5,5],[5,6],[5,7],[6,7],[7,7],[7,6],[6,6],[6,5],[7,5],[8,5],[9,5],[10,5],[10,4],[10,3],[11,3],[11,4],[11,5],[12,5],[13,5],[14,5],[15,5],[16,5],[16,6],[16,7],[17,7],[18,7],[18,8],[18,9],[18,10],[18,11],[18,12],[18,13],[18,14],[18,15],[18,16],[17,16],[16,16],[15,16],[15,17],[16,17],[17,17],[18,17],[18,18],[18,19],[19,19],[19,18],[19,17],[19,16],[19,15],[19,14],[19,13],[19,12],[19,11],[19,10],[19,9],[19,8],[19,7],[19,6],[19,5],[19,4],[19,3],[19,2],[18,2],[18,3],[18,4],[18,5],[18,6],[17,6],[17,5],[17,4],[17,3],[17,2],[17,1],[18,1],[18,0],[17,0],[16,0],[15,0]]}],"height":20,"game_id":"01d41d8a-ae1c-4334-9c2c-0e0ee9dcbb76","food":[[15,9],[15,12],[5,1],[0,17],[13,9],[4,19]],"dead_snakes":[{"taunt":"I'm a snake, not a doctor","name":"Genetisnake NBK","id":"77bae9cb-e95f-42be-9bfc-e94d500f5b48","health_points":99,"coords":[[7,18],[7,19],[7,18]]}]}`

const gameString16 = `{"you":"f9dc8c4c-a918-4564-a901-f2dfc2b9b2df","width":20,"turn":14,"snakes":[{"taunt":"","name":"Type Snake","id":"541457d9-252b-41b7-a48c-d49921bb876d","health_points":89,"coords":[[12,12],[12,13],[12,14],[12,15],[12,16]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"f9dc8c4c-a918-4564-a901-f2dfc2b9b2df","health_points":93,"coords":[[11,11],[11,12],[11,13],[11,14]]},{"taunt":"I'm a snake, not a doctor","name":"Genetisnake NBK","id":"c234d8a5-e876-48f6-99df-ae58be8c3bc1","health_points":86,"coords":[[4,19],[4,18],[5,18]]},{"taunt":"hungry hungry snake","name":"NBK Greedy","id":"7b8f820e-1a1f-446b-97cd-bc01198a15e5","health_points":95,"coords":[[1,1],[1,2],[1,3],[1,4],[1,5],[1,6]]}],"height":20,"game_id":"66e8b0e2-4ae3-48d3-bf22-156ffecee5e2","food":[[9,5],[18,4],[0,11],[2,19],[12,8],[1,0],[0,18],[13,0],[5,3],[2,14]],"dead_snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"8590987f-fe28-44eb-88e7-52e97c041967","health_points":99,"coords":[[4,1],[4,2],[4,1]]}]}`

const gameString17 = `{"you":"8dffcb83-e48c-4795-8406-9976ca317ae8","width":20,"turn":53,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"d2f4f1b0-e2a4-40ea-bbff-ffc8a9ab8e50","health_points":99,"coords":[[0,1],[0,0],[1,0],[2,0],[3,0],[4,0],[5,0],[6,0],[6,1],[6,2],[6,3],[6,4],[6,5]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"f55c3d87-4ad3-4e45-b3d6-084d3e1b9adc","health_points":100,"coords":[[1,12],[2,12],[2,11],[2,10],[3,10],[4,10],[5,10],[6,10],[7,10],[7,11],[6,11],[6,12],[7,12],[8,12],[9,12],[9,12]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"6f8924f1-0da7-4b84-81e3-8c433d324a27","health_points":98,"coords":[[4,15],[4,14],[5,14],[5,15],[5,16],[5,17],[5,18],[4,18],[3,18],[3,17],[4,17],[4,16],[3,16],[2,16],[1,16],[1,17]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"fc9b848a-14d1-4964-9c9c-d0829840b0ae","health_points":74,"coords":[[10,1],[10,2],[10,3],[10,4],[10,5],[9,5],[8,5],[8,6],[9,6]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"a92591d6-10ec-4752-81d8-fbe2d3bfc5f4","health_points":99,"coords":[[13,9],[14,9],[14,10],[14,11],[14,12],[14,13],[14,14],[13,14],[13,15],[12,15]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"8dffcb83-e48c-4795-8406-9976ca317ae8","health_points":84,"coords":[[9,0],[9,1],[9,2],[9,3],[8,3],[8,2],[8,1],[8,0]]}],"height":20,"game_id":"66e8b0e2-4ae3-48d3-bf22-156ffecee5e2","food":[[16,19],[11,4],[7,19],[0,19],[15,1],[8,10],[16,4],[17,3],[3,4],[5,4]],"dead_snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"6300cdba-2fb9-42d2-ac97-b7504cceb748","health_points":87,"coords":[[8,-1],[8,0],[7,0]]}]}`
