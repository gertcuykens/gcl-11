import com.google.appengine.api.datastore.*;

import java.io.IOException;
import java.io.PrintWriter;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class DataStoreServlet extends HttpServlet {
    @Override
    public void doGet(HttpServletRequest req, HttpServletResponse resp) throws IOException {
        try {
            DatastoreService datastore = DatastoreServiceFactory.getDatastoreService();
            Entity userEntity = datastore.get(KeyFactory.createKey("UserInfo", "test"));
            Query query = new Query("Task");
            query.addFilter("dueDate", Query.FilterOperator.LESS_THAN, "test");
            for (Entity taskEntity : datastore.prepare(query).asIterable()) {
                if ("done".equals(taskEntity.getProperty("status"))) {
                    datastore.delete(taskEntity.getKey());
                } else {
                    taskEntity.setProperty("status", "overdue");
                    datastore.put(taskEntity);
                }
            }
        } catch (Exception e) {
            resp.setContentType("text/plain");
            PrintWriter out = resp.getWriter();
            out.print(e.toString());
        }
    }
}
