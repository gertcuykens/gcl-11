package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.*;
import android.view.View.OnClickListener;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

public class MainActivity extends Activity {
    public static Activity activity;
    EditText graphValue;
    EditText messageValue;
    Button submitButton;
    Button publishButton;
    TextView textViewResults;

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.gcl11);
        activity = this;
        
        this.messageValue = (EditText) findViewById(R.id.messageValue);
        this.graphValue = (EditText) findViewById(R.id.graphValue);

        this.submitButton = (Button) findViewById(R.id.submitButton);
        this.submitButton.setOnClickListener(new OnClickListener() {
            public void onClick(View view) {onClickSubmit(view);}
        });

        this.publishButton = (Button) findViewById(R.id.publishButton);
        this.publishButton.setOnClickListener(new OnClickListener() {
            public void onClick(View view) {onClickRequest(view);}
        });

        this.textViewResults = (TextView) findViewById(R.id.textViewResults);
        this.textViewResults.setOnClickListener(new OnClickListener() {
            public void onClick(View view) {
                onClickList(view);
            }
        });

        Intent intent = new Intent(this, LoginUsingLoginFragmentActivity.class);
        startActivity(intent);
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        MenuInflater inflater = getMenuInflater();
        inflater.inflate(R.menu.login, menu);
        return super.onCreateOptionsMenu(menu);
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case R.id.action_settings:
                Intent intent = new Intent(this, LoginUsingLoginFragmentActivity.class);
                startActivity(intent);
                return true;
            default:
                return super.onOptionsItemSelected(item);
        }
    }

    private void onClickSubmit(View view) {
        Global g = Global.getInstance();
        g.setMessage(messageValue.getText().toString());
        Context context = view.getContext();
        new SubmitTask().execute(context);
    }

    private void onClickList(View view) {
        Context context = view.getContext();
        new ListTask().execute(context);
    }

    private void onClickRequest(View view) {
        Global g = Global.getInstance();
        g.setMessage(messageValue.getText().toString());
        g.setGraph(graphValue.getText().toString());
        Context context = view.getContext();
        new RequestTask().execute(context);
    }

    @Override
    public void onActivityResult(int requestCode, int resultCode, Intent data) {
        Log.i("graph", "----------------------------");
        Log.i("graph", String.valueOf("MainActivity onActivityResult:"+resultCode));
        Log.i("graph", "----------------------------");
    }

    static void toaster(Context c, String s) {
        Toast toast = Toast.makeText(c, s, Toast.LENGTH_SHORT);
        toast.setGravity(Gravity.TOP|Gravity.RIGHT, 0, 0);
        toast.show();
    }

}

/*
    @Override
    public void onActivityResult(int requestCode, int resultCode, Intent data) {
        new ListTask().execute(this);
    }

    static void alert(Context c, String s) {
        new AlertDialog.Builder(c)
                .setTitle(R.string.login_failed_dialog_title)
                .setMessage(s)
                .setPositiveButton(R.string.ok_button, null)
                .show();
    }
*/

//import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
//import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;

/*
   private void sendRequests() {

        Bundle postParams = new Bundle();
        postParams.putString("message",messageValue.getText().toString());
        postParams.putString("name","");
        postParams.putString("link","");
        postParams.putString("picture","");
        //postParams.putString("access_token", "");

        textViewResults.setText(String.valueOf(postParams)+"\n");

        Request request =new Request(session, "me", null, null, new Request.Callback() {
            public void onCompleted(Response response) {
                GraphObject graphObject = response.getGraphObject();
                FacebookRequestError error = response.getError();
                String s = textViewResults.getText().toString();

                if (graphObject != null) {
                    if (graphObject.getProperty("id") != null) {
                        s = s + String.format("%s: %s\n", graphObject.getProperty("id"), graphObject.getProperty("name"));
                    } else {
                        s = s + String.format("%s: <no such id>\n", "me");
                    }
                } else if (error != null) {
                    s = s + String.format("Error: %s", error.getErrorMessage());
                }

                textViewResults.setText(s);
            }
        });

        Request.Callback callback = new Request.Callback() {
            public void onCompleted(Response response) {
                GraphObject graphObject = response.getGraphObject();
                FacebookRequestError error = response.getError();
                String s = textViewResults.getText().toString();

                if (error != null) {
                    s = s + String.format("Error gcl-11: %s\n", error.getErrorMessage());
                    Log.e("FACEBOOK ERROR", ""+ error.getErrorMessage());
                } else {
                    JSONObject graphResponse = graphObject.getInnerJSONObject();
                    try {
                        String postId = graphResponse.getString("id");
                        s = s + String.format("%s: gcl-11\n", postId);
                        Log.i ("graph", postId);
                    } catch (JSONException e) {
                        s = s + String.format("Error gcl-11: %s\n", e.getMessage());
                        Log.e("FACEBOOK ERROR", e.getMessage());
                    }
                }

                textViewResults.setText(s);
            }
        };

        pendingRequest = false;

        request.executeAsync();
        Request request2 = new Request(session, graphValue.getText().toString(), postParams, HttpMethod.POST, callback);
        RequestAsyncTask task = new RequestAsyncTask(request2);
        task.execute();

    }

    private void sendRequests() {
        textViewResults.setText("");

        String requestIdsText = messageValue.getText().toString();
        String[] requestIds = requestIdsText.split(",");

        List<Request> requests = new ArrayList<Request>();
        for (final String requestId : requestIds) {
            requests.add(new Request(session, requestId, null, null, new Request.Callback() {
                public void onCompleted(Response response) {
                    GraphObject graphObject = response.getGraphObject();
                    FacebookRequestError error = response.getError();
                    String s = textViewResults.getText().toString();
                    if (graphObject != null) {
                        if (graphObject.getProperty("id") != null) {
                            s = s + String.format("%s: %s\n", graphObject.getProperty("id"), graphObject.getProperty(
                                    "name"));
                        } else {
                            s = s + String.format("%s: <no such id>\n", requestId);
                        }
                    } else if (error != null) {
                        s = s + String.format("Error: %s", error.getErrorMessage());
                    }
                    textViewResults.setText(s);
                }
            }));
        }
        pendingRequest = false;
        Request.executeBatchAsync(requests);
    }
*/

