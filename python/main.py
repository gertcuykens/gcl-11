from google.appengine.api import users
from google.appengine.ext import ndb
import webapp2

class MainPage(webapp2.RequestHandler):
    def get(self):
        user = users.get_current_user()
        if not user:
            self.redirect(users.create_login_url(self.request.uri))
        else:
            self.response.headers['Content-Type'] = 'text/plain'
            self.response.out.write('Hello, '+ user.nickname())

app = webapp2.WSGIApplication([('/', MainPage)], debug=True)
