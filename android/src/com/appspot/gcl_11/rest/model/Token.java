/*
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */
/*
 * This code was generated by https://code.google.com/p/google-apis-client-generator/
 * (build: 2014-04-15 19:10:39 UTC)
 * on 2014-04-26 at 20:35:14 UTC 
 * Modify at your own risk.
 */

package com.appspot.gcl_11.rest.model;

/**
 * Model definition for Token.
 *
 * <p> This is the Java data model class that specifies how to parse/serialize into the JSON that is
 * transmitted over HTTP when working with the rest. For a detailed explanation see:
 * <a href="http://code.google.com/p/google-http-java-client/wiki/JSON">http://code.google.com/p/google-http-java-client/wiki/JSON</a>
 * </p>
 *
 * @author Google, Inc.
 */
@SuppressWarnings("javadoc")
public final class Token extends com.google.api.client.json.GenericJson {

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("access_token")
  private java.lang.String accessToken;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("email_token")
  private java.lang.String emailToken;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("expires_in") @com.google.api.client.json.JsonString
  private java.lang.Long expiresIn;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key
  private com.google.api.client.util.DateTime expiry;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key
  private java.util.List<Property> extra;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("id_token") @com.google.api.client.json.JsonString
  private java.lang.Long idToken;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key
  private java.lang.String message;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("name_token")
  private java.lang.String nameToken;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("oauth_token")
  private java.lang.String oauthToken;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("oauth_verifier")
  private java.lang.String oauthVerifier;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("refresh_token")
  private java.lang.String refreshToken;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key
  private java.lang.String status;

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("type_token")
  private java.lang.String typeToken;

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getAccessToken() {
    return accessToken;
  }

  /**
   * @param accessToken accessToken or {@code null} for none
   */
  public Token setAccessToken(java.lang.String accessToken) {
    this.accessToken = accessToken;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getEmailToken() {
    return emailToken;
  }

  /**
   * @param emailToken emailToken or {@code null} for none
   */
  public Token setEmailToken(java.lang.String emailToken) {
    this.emailToken = emailToken;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.Long getExpiresIn() {
    return expiresIn;
  }

  /**
   * @param expiresIn expiresIn or {@code null} for none
   */
  public Token setExpiresIn(java.lang.Long expiresIn) {
    this.expiresIn = expiresIn;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public com.google.api.client.util.DateTime getExpiry() {
    return expiry;
  }

  /**
   * @param expiry expiry or {@code null} for none
   */
  public Token setExpiry(com.google.api.client.util.DateTime expiry) {
    this.expiry = expiry;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.util.List<Property> getExtra() {
    return extra;
  }

  /**
   * @param extra extra or {@code null} for none
   */
  public Token setExtra(java.util.List<Property> extra) {
    this.extra = extra;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.Long getIdToken() {
    return idToken;
  }

  /**
   * @param idToken idToken or {@code null} for none
   */
  public Token setIdToken(java.lang.Long idToken) {
    this.idToken = idToken;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getMessage() {
    return message;
  }

  /**
   * @param message message or {@code null} for none
   */
  public Token setMessage(java.lang.String message) {
    this.message = message;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getNameToken() {
    return nameToken;
  }

  /**
   * @param nameToken nameToken or {@code null} for none
   */
  public Token setNameToken(java.lang.String nameToken) {
    this.nameToken = nameToken;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getOauthToken() {
    return oauthToken;
  }

  /**
   * @param oauthToken oauthToken or {@code null} for none
   */
  public Token setOauthToken(java.lang.String oauthToken) {
    this.oauthToken = oauthToken;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getOauthVerifier() {
    return oauthVerifier;
  }

  /**
   * @param oauthVerifier oauthVerifier or {@code null} for none
   */
  public Token setOauthVerifier(java.lang.String oauthVerifier) {
    this.oauthVerifier = oauthVerifier;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getRefreshToken() {
    return refreshToken;
  }

  /**
   * @param refreshToken refreshToken or {@code null} for none
   */
  public Token setRefreshToken(java.lang.String refreshToken) {
    this.refreshToken = refreshToken;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getStatus() {
    return status;
  }

  /**
   * @param status status or {@code null} for none
   */
  public Token setStatus(java.lang.String status) {
    this.status = status;
    return this;
  }

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getTypeToken() {
    return typeToken;
  }

  /**
   * @param typeToken typeToken or {@code null} for none
   */
  public Token setTypeToken(java.lang.String typeToken) {
    this.typeToken = typeToken;
    return this;
  }

  @Override
  public Token set(String fieldName, Object value) {
    return (Token) super.set(fieldName, value);
  }

  @Override
  public Token clone() {
    return (Token) super.clone();
  }

}
