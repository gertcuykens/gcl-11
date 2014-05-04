/**
 * Copyright 2010-present Facebook.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package my.endpoints;

import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.support.v4.app.FragmentActivity;
import android.support.v4.app.FragmentManager;
import android.util.Log;
import com.facebook.LoggingBehavior;
import com.facebook.Session;
import com.facebook.SessionState;
import com.facebook.Settings;
import com.facebook.widget.UserSettingsFragment;

public class LoginUsingLoginFragmentActivity extends FragmentActivity {
    private UserSettingsFragment userSettingsFragment;
    private Context context;

    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        context = MainActivity.activity;
        setContentView(R.layout.login_fragment_activity);

        Settings.addLoggingBehavior(LoggingBehavior.INCLUDE_ACCESS_TOKENS);

        FragmentManager fragmentManager = getSupportFragmentManager();
        userSettingsFragment = (UserSettingsFragment) fragmentManager.findFragmentById(R.id.login_fragment);
        userSettingsFragment.setSessionStatusCallback(new Session.StatusCallback() {
            @Override
            public void call(Session session, SessionState state, Exception exception) {
                Log.i("graph", "----------------------------");
                Log.i("graph", String.format("New session state: %s", state.toString()));
                Log.i("graph", "----------------------------");
                if (state.isOpened()){new ListTask().execute(context);}
            }
        });
    }

    @Override
    public void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        userSettingsFragment.onActivityResult(requestCode, resultCode, data);
        Log.i("graph", "----------------------------");
        Log.i("graph", String.valueOf("LoginUsingLoginFragmentActivity onActivityResult: "+resultCode));
        Log.i("graph", "----------------------------");
    }

}
