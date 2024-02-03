conn = new Mongo();
db = conn.getDB("StreamFlow");

// Add a test user.
user = { "username": "test@costreamflows.com", "diplay_name": "test_user", "password": "1234" };
insertUserResult = db.Users.insert(user);
